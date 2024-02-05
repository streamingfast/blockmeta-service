package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pbbmsrv "github.com/streamingfast/blockmeta-service/pb/sf/blockmeta/v2"
	"github.com/streamingfast/blockmeta-service/service"
	"google.golang.org/grpc"
)

var (
	listenAddress     = flag.String("grpc-listen-addr ", ":9000", "The server port")
	sinkServerAddress = flag.String("grpc-listen-addr ", "", "The server port")
)

// todo: convert to cobra and viper
func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *listenAddress))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sinkClient := service.ConnectToSinkServer(*sinkServerAddress)
	blockService := service.NewBlockService(sinkClient)
	blockByTimeService := service.NewBlockByTimeService(sinkClient)

	sBlock := grpc.NewServer()
	pbbmsrv.RegisterBlockServer(sBlock, blockService)
	pbbmsrv.RegisterBlockByTimeServer(sBlock, blockByTimeService)

	log.Printf("server block listening at %v", listener.Addr())
	if err := sBlock.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
