package main

import (
	"context"
	"flag"
	"os"
	"strings"
	"time"

	"github.com/streamingfast/dgrpc/server/factory"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pbbmsrv "github.com/streamingfast/blockmeta-service/pb/sf/blockmeta/v2"
	"github.com/streamingfast/blockmeta-service/service"
	derr "github.com/streamingfast/derr"
	dgrpcserver "github.com/streamingfast/dgrpc/server"
)

var (
	listenAddress     = flag.String("grpc-listen-addr", "", "The gRPC server listen address")
	sinkServerAddress = flag.String("sink-addr", "", "The sink server address")
)

func main() {
	flag.Parse()

	if *sinkServerAddress == "" {
		logger.Error("sink server address is required")
		os.Exit(1)
	}

	if *listenAddress == "" {
		logger.Error("listen address is required")
		os.Exit(1)
	}

	sinkClient := service.ConnectToSinkServer(*sinkServerAddress)
	blockService := service.NewBlockService(sinkClient)
	blockByTimeService := service.NewBlockByTimeService(sinkClient)

	options := []dgrpcserver.Option{
		dgrpcserver.WithLogger(zap.NewNop()),
		dgrpcserver.WithHealthCheck(dgrpcserver.HealthCheckOverGRPC|dgrpcserver.HealthCheckOverHTTP, healthCheck()),
	}

	if strings.Contains(*listenAddress, "*") {
		options = append(options, dgrpcserver.WithInsecureServer())
	} else {
		options = append(options, dgrpcserver.WithPlainTextServer())
	}

	cleanListenAddress := strings.ReplaceAll(*listenAddress, "*", "")

	grpcServer := factory.ServerFromOptions(options...)
	grpcServer.RegisterService(func(gs grpc.ServiceRegistrar) {
		pbbmsrv.RegisterBlockServer(gs, blockService)
		pbbmsrv.RegisterBlockByTimeServer(gs, blockByTimeService)
	})

	go func() {
		logger.Info("launching gRPC server", "listen_address", cleanListenAddress)
		grpcServer.Launch(cleanListenAddress)
	}()

	<-derr.SetupSignalHandler(30 * time.Second)
}

func healthCheck() dgrpcserver.HealthCheck {
	return func(ctx context.Context) (isReady bool, out interface{}, err error) {
		if derr.IsShuttingDown() {
			return false, nil, nil
		}
		return true, nil, nil
	}
}
