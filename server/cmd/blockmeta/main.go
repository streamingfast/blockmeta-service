package main

import (
	"flag"
	"os"
	"regexp"
	"time"

	server "github.com/streamingfast/blockmeta-service/server"
	pbbmsrv "github.com/streamingfast/blockmeta-service/server/pb/sf/blockmeta/v2"
	"github.com/streamingfast/blockmeta-service/server/service"
	"github.com/streamingfast/dauth"
	"github.com/streamingfast/derr"
	"github.com/streamingfast/dgrpc/server/factory"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	listenAddress          = flag.String("grpc-listen-addr", "", "The gRPC server listen address")
	sinkServerAddress      = flag.String("sink-addr", "", "The sink server address")
	authUrl                = flag.String("auth-url", "null://", "The URL of the auth server")
	corsHostRegexAllowFlag = flag.String("cors-host-regex-allow", "^localhost", "Regex to allow CORS origin requests from, defaults to localhost only")
)

func main() {
	flag.Parse()

	if *sinkServerAddress == "" {
		zlog.Error("sink server address is required")
		os.Exit(1)
	}

	if *listenAddress == "" {
		zlog.Error("listen address is required")
		os.Exit(1)
	}

	sinkClient := server.ConnectToSinkServer(*sinkServerAddress)
	blockService := service.NewBlockService(sinkClient)
	blockByTimeService := service.NewBlockByTimeService(sinkClient)

	auth, err := dauth.New(*authUrl, zlog)
	if err != nil {
		zlog.Error("unable to create authenticator", zap.Error(err))
		os.Exit(1)
	}

	var corsHostRegexAllow *regexp.Regexp
	if *corsHostRegexAllowFlag != "" {
		hostRegex, err := regexp.Compile(*corsHostRegexAllowFlag)
		if err != nil {
			zlog.Error("unable to compile cors-host-regex-allow", zap.Error(err))
			os.Exit(1)
		}
		corsHostRegexAllow = hostRegex
	}

	grpcServer := server.NewGrpcServer(corsHostRegexAllow, zlog, *listenAddress, auth)

	grpcServer := factory.ServerFromOptions(options...)
	grpcServer.RegisterService(func(gs grpc.ServiceRegistrar) {
		pbbmsrv.RegisterBlockServer(gs, blockService)
		pbbmsrv.RegisterBlockByTimeServer(gs, blockByTimeService)
	})

	go func() {
		zlog.Info("launching gRPC server", zap.String("listen_address", cleanListenAddress))
		grpcServer.Launch(cleanListenAddress)
	}()

	<-derr.SetupSignalHandler(30 * time.Second)
}
