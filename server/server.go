package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"github.com/rs/cors"
	pbbmsrv "github.com/streamingfast/blockmeta-service/server/pb/sf/blockmeta/v2"
	"github.com/streamingfast/blockmeta-service/server/pb/sf/blockmeta/v2/pbbmsrvconnect"
	"github.com/streamingfast/dauth"
	dauthconnect "github.com/streamingfast/dauth/middleware/connect"
	"github.com/streamingfast/derr"
	dgrpcserver "github.com/streamingfast/dgrpc/server"
	connectweb "github.com/streamingfast/dgrpc/server/connect-web"
	"github.com/streamingfast/shutter"
	pbkv "github.com/streamingfast/substreams-sink-kv/pb/substreams/sink/kv/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	*shutter.Shutter
	shutdownLock sync.RWMutex

	corsHostRegexAllow *regexp.Regexp
	logger             *zap.Logger
	authenticator      dauth.Authenticator
	httpListenAddr     string
	sinkClient         pbkv.KvClient
}

func NewGrpcServer(corsHostRegexAllow *regexp.Regexp, logger *zap.Logger, httpListenAddr string, authenticator dauth.Authenticator, sinkClient pbkv.KvClient) *GrpcServer {
	return &GrpcServer{
		corsHostRegexAllow: corsHostRegexAllow,
		logger:             logger,
		authenticator:      authenticator,
		httpListenAddr:     httpListenAddr,
		sinkClient:         sinkClient,
	}
}

func (s *GrpcServer) Run(ctx context.Context) {
	s.logger.Debug("starting server")

	options := []dgrpcserver.Option{
		dgrpcserver.WithLogger(s.logger),
		dgrpcserver.WithHealthCheck(dgrpcserver.HealthCheckOverGRPC|dgrpcserver.HealthCheckOverHTTP, s.healthCheck()),
		dgrpcserver.WithConnectInterceptor(dauthconnect.NewAuthInterceptor(s.authenticator, s.logger)),
		dgrpcserver.WithReflection(pbbmsrv.Block_ServiceDesc.ServiceName),
		dgrpcserver.WithReflection(pbbmsrv.BlockByTime_ServiceDesc.ServiceName),
		dgrpcserver.WithCORS(s.corsOption()),
	}

	if strings.Contains(s.httpListenAddr, "*") {
		s.logger.Warn("grpc server with insecure server")
		options = append(options, dgrpcserver.WithInsecureServer())
	} else {
		s.logger.Info("grpc server with plain text server")
		options = append(options, dgrpcserver.WithPlainTextServer())
	}

	streamHandlerGetter := func(opts ...connect.HandlerOption) (string, http.Handler) {
		return pbbmsrvconnect.NewBlockHandler(s, opts...)

	}
	srv := connectweb.New([]connectweb.HandlerGetter{streamHandlerGetter}, options...)
	addr := strings.ReplaceAll(s.httpListenAddr, "*", "")

	s.OnTerminating(func(err error) {
		s.shutdownLock.Lock()
		s.logger.Warn("shutting down connect web server")

		time.Sleep(1 * time.Second)

		srv.Shutdown(nil)
		s.logger.Warn("connect web server shutdown")
	})

	s.OnTerminated(func(err error) {
		s.shutdownLock.Unlock()
	})

	srv.Launch(addr)
	<-srv.Terminated()
}

func (s *GrpcServer) corsOption() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: s.allowedOrigin,
		AllowedHeaders:  []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-State",
			"Grpc-State-Details-Bin",
		},
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

func (s *GrpcServer) allowedOrigin(origin string) bool {
	s.logger.Debug("allowed origin", zap.String("origin", origin))

	if s.corsHostRegexAllow == nil {
		s.logger.Warn("allowed origin, no host regex allowed filter specify denying origin", zap.String("origin", origin))
		return false
	}

	uri, err := url.Parse(origin)
	if err != nil {
		s.logger.Warn("failed to parse origin", zap.String("origin", origin), zap.Error(err))
		return false
	}
	return s.corsHostRegexAllow.MatchString(uri.Host)
}

func (s *GrpcServer) healthCheck() dgrpcserver.HealthCheck {
	return func(ctx context.Context) (isReady bool, out interface{}, err error) {
		if derr.IsShuttingDown() {
			return false, nil, nil
		}
		return true, nil, nil
	}
}

func (s *GrpcServer) NumToID(ctx context.Context, in *connect.Request[pbbmsrv.NumToIDReq]) (*connect.Response[pbbmsrv.BlockResp], error) {
	slog.Info("handling NumToID request", "block_num", in.Msg.BlockNum)
	prefix := Keyer.PackNumPrefixKey(in.Msg.BlockNum)

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		//grpcError already handled
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, status.Errorf(13, "more than one block found for block number: %v", in.Msg.BlockNum)
	}

	blockNum, blockID, err := Keyer.UnpackNumIDKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, status.Errorf(13, "error unpacking block number and block ID: %v", err)
	}

	blockPbTimestamp := valueToTimestamp(response.KeyValues[0].Value)
	return &connect.Response[pbbmsrv.BlockResp]{Msg: &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}}, nil
}

func (s *GrpcServer) IDToNum(ctx context.Context, in *connect.Request[pbbmsrv.IDToNumReq]) (*connect.Response[pbbmsrv.BlockResp], error) {

	slog.Info("handling IDToNum request", "block_id", in.Msg.BlockID)
	prefix := Keyer.PackIDPrefixKey(in.Msg.BlockID)

	if prefix == "1:" {
		return nil, status.Errorf(3, "block id is empty")
	}
	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, status.Errorf(13, "more than one block found for block id: %v", in.Msg.BlockID)
	}

	blockNum, blockID, err := Keyer.UnpackIDNumKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, status.Errorf(13, "error unpacking block number and block ID: %v", err)
	}

	blockPbTimestamp := valueToTimestamp(response.KeyValues[0].Value)
	return &connect.Response[pbbmsrv.BlockResp]{Msg: &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}}, nil
}

func (s *GrpcServer) Head(ctx context.Context, _ *connect.Request[pbbmsrv.Empty]) (*connect.Response[pbbmsrv.BlockResp], error) {
	slog.Info("handling Head request")
	prefix := TblPrefixTimelineBck + ":"

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix, Limit: 1})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	blockPbTimestamp, blockID, err := Keyer.UnpackTimeIDKey(response.KeyValues[0].Key, false)
	if err != nil {
		return nil, status.Errorf(13, "error unpacking block number and block ID: %v", err)
	}

	blockNum := valueToBlockNumber(response.KeyValues[0].Value)

	return &connect.Response[pbbmsrv.BlockResp]{Msg: &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}}, nil
}

func (s *GrpcServer) At(ctx context.Context, in connect.Request[pbbmsrv.TimeReq]) (*connect.Response[pbbmsrv.BlockResp], error) {
	slog.Info("handling At request", "block_time", in.Msg.Time)
	prefix := Keyer.PackTimePrefixKey(in.Msg.Time.AsTime(), false)

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, status.Errorf(500, "more than one block found for block timestamp: %v", in.Msg.Time)
	}

	blockPbTimestamp, blockID, err := Keyer.UnpackTimeIDKey(response.KeyValues[0].Key, false)
	if err != nil {
		return nil, status.Errorf(500, "error unpacking block number and block ID: %w", err)
	}

	blockNum := valueToBlockNumber(response.KeyValues[0].Value)
	return &connect.Response[pbbmsrv.BlockResp]{Msg: &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}}, nil
}

func (s *GrpcServer) Before(ctx context.Context, in *connect.Request[pbbmsrv.RelativeTimeReq]) (*connect.Response[pbbmsrv.BlockResp], error) {
	slog.Info("handling Before request", "block_time", in.Msg.Time)
	prefix := Keyer.PackTimePrefixKey(in.Msg.Time.AsTime(), false)

	response, err := s.sinkClient.Scan(ctx, &pbkv.ScanRequest{Begin: prefix, Limit: 4})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	var blockID string
	var blockNum uint64
	blockPbTimestamp := &timestamppb.Timestamp{}

	for i := 0; i < len(response.KeyValues); i++ {
		blockPbTimestamp, blockID, err = Keyer.UnpackTimeIDKey(response.KeyValues[i].Key, false)
		if err != nil {
			return nil, status.Errorf(13, "error unpacking block number and block ID: %w", err)
		}

		if !in.Msg.Inclusive && (blockPbTimestamp.AsTime() == in.Msg.Time.AsTime()) {
			continue
		}

		blockNum = valueToBlockNumber(response.KeyValues[i].Value)
		break
	}
	return &connect.Response[pbbmsrv.BlockResp]{Msg: &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}}, nil
}

func (s *GrpcServer) After(ctx context.Context, in *pbbmsrv.RelativeTimeReq) (*pbbmsrv.BlockResp, error) {
	slog.Info("handling After request", "block_time", in.Time)
	prefix := Keyer.PackTimePrefixKey(in.Time.AsTime(), true)

	response, err := s.sinkClient.Scan(ctx, &pbkv.ScanRequest{Begin: prefix, Limit: 4})
	if err != nil {
		return nil, status.Errorf(13, "error getting block data from sink server: %w", err)
	}

	var blockID string
	var blockNum uint64
	blockPbTimestamp := &timestamppb.Timestamp{}

	for i := 0; i < len(response.KeyValues); i++ {

		blockPbTimestamp, blockID, err = Keyer.UnpackTimeIDKey(response.KeyValues[i].Key, true)
		if err != nil {
			return nil, status.Errorf(13, "error unpacking block number and block ID: %w", err)
		}

		if !in.Inclusive && (blockPbTimestamp.AsTime() == in.Time.AsTime()) {
			continue
		}

		blockNum = valueToBlockNumber(response.KeyValues[i].Value)

		break
	}
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}
