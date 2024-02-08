package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	pbbmsrv "github.com/streamingfast/blockmeta-service/pb/sf/blockmeta/v2"
	"github.com/streamingfast/blockmeta-service/service"
	"google.golang.org/grpc"
)

var (
	listenAddress     = flag.String("grpc-listen-addr", ":9000", "The gRPC server listen address")
	sinkServerAddress = flag.String("sink-addr", "", "The sink server address")
)

// todo: convert to cobra and viper
func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s", *listenAddress))
	if err != nil {
		logger.Error("failed to listen", "error", err, "address", *listenAddress)
		os.Exit(1)
	}

	if *sinkServerAddress == "" {
		logger.Error("sink server address is required")
		os.Exit(1)
	}

	sinkClient := service.ConnectToSinkServer(*sinkServerAddress)
	blockService := service.NewBlockService(sinkClient)
	blockByTimeService := service.NewBlockByTimeService(sinkClient)

	sBlock := grpc.NewServer()
	pbbmsrv.RegisterBlockServer(sBlock, blockService)
	pbbmsrv.RegisterBlockByTimeServer(sBlock, blockByTimeService)

	logger.Info("server block listening", "address", listener.Addr())
	if err := sBlock.Serve(listener); err != nil {
		logger.Error("failed to serve", "error", err)
		os.Exit(1)
	}
}
