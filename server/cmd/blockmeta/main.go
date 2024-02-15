package main

import (
	"context"
	"flag"
	"os"
	"regexp"

	"github.com/streamingfast/blockmeta-service/server"
	"github.com/streamingfast/dauth"
	authGRPC "github.com/streamingfast/dauth/grpc"
	authNull "github.com/streamingfast/dauth/null"
	"github.com/streamingfast/derr"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var (
	listenAddress          = flag.String("grpc-listen-addr", ":9000", "The gRPC server listen address")
	sinkServerAddress      = flag.String("sink-addr", ":9001", "The sink server address")
	authUrl                = flag.String("auth-url", "null://", "The URL of the auth server")
	corsHostRegexAllowFlag = flag.String("cors-host-regex-allow", "^localhost", "Regex to allow CORS origin requests from, defaults to localhost only")
)

var zlog, _ = logging.ApplicationLogger("blockmeta", "github.com/streamingfast/blockmeta-service/server/cmd/blockmeta")

func main() {
	flag.Parse()
	ctx := context.Background()

	authNull.Register()
	authGRPC.Register()

	if *sinkServerAddress == "" {
		zlog.Error("sink server address is required")
		os.Exit(1)
	}

	if *listenAddress == "" {
		zlog.Error("listen address is required")
		os.Exit(1)
	}

	sinkClient := server.ConnectToSinkServer(*sinkServerAddress)

	authenticator, err := dauth.New(*authUrl, zlog)
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

	grpcServer := server.NewGrpcServer(*listenAddress, sinkClient, corsHostRegexAllow, authenticator, zlog)
	signal := derr.SetupSignalHandler(0)
	go func() {
		<-signal
		grpcServer.Shutdown(nil)
	}()

	grpcServer.Run(ctx)
	<-grpcServer.Terminated()
	if grpcServer.Err() != nil {
		zlog.Error("server terminated with error", zap.Error(grpcServer.Err()))
		os.Exit(1)
	}
}
