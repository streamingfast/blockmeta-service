package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pbbmsrv "github.com/streamingfast/blockmeta-service/server/pb/sf/blockmeta/v2"
	"github.com/streamingfast/dgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	serverAddress = flag.String("server-addr", "localhost:50051", "The server address")
)

func main() {
	flag.Parse()
	if *serverAddress == "" {
		log.Fatalf("You must provide a server address using the -server-addr flag.")
	}

	conn, err := dgrpc.NewInternalClient(*serverAddress)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pbbmsrv.NewBlockClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	blockResp, err := c.NumToID(ctx, &pbbmsrv.NumToIDReq{BlockNum: 6})
	if err != nil {
		log.Fatalf("could not get block data  NumToID request: %v", err)
	}

	blockSecondResp, err := c.IDToNum(ctx, &pbbmsrv.IDToNumReq{BlockID: "f37c632d361e0a93f08ba29b1a2c708d9caa3ee19d1ee8d2a02612bffe49f0a9"})
	if err != nil {
		log.Fatalf("could not get block data for IDToNum request: %v", err)
	}

	d := pbbmsrv.NewBlockByTimeClient(conn)
	blockThirdResp, err := d.At(ctx, &pbbmsrv.TimeReq{Time: timestamppb.New(time.UnixMilli(1438270241000))})
	if err != nil {
		log.Fatalf("could not get block data for At request: %v", err)
	}

	blockFourthResp, err := d.After(ctx, &pbbmsrv.RelativeTimeReq{Time: timestamppb.New(time.UnixMilli(1438270241000)), Inclusive: true})
	if err != nil {
		log.Fatalf("could not get block data for After request: %v", err)
	}

	blockFifthResp, err := d.Before(ctx, &pbbmsrv.RelativeTimeReq{Time: timestamppb.New(time.UnixMilli(1438270241000)), Inclusive: false})
	if err != nil {
		log.Fatalf("could not get block data for Before request: %v", err)
	}

	blockSithResp, err := c.Head(ctx, &pbbmsrv.Empty{})

	fmt.Printf("Blockresponse for NumToID request: %s\n", blockResp.String())

	fmt.Printf("Blockresponse for  IDToNum request: %s\n", blockSecondResp.String())

	fmt.Printf("Blockresponse for  At request: %s\n", blockThirdResp.String())

	fmt.Printf("Blockresponse for After request: %s\n", blockFourthResp.String())

	fmt.Printf("Blockresponse for Before request: %s\n", blockFifthResp.String())

	fmt.Printf("Blockresponse for Head request: %s\n", blockSithResp.String())

}
