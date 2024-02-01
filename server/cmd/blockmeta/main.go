package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"

	blockmeta "blockmeta_server/pb/blockmeta"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type serverBlock struct {
	blockmeta.UnimplementedBlockServer
	blockmeta.UnimplementedBlockByTimeServer
}

func (s *serverBlock) NumToID(ctx context.Context, in *blockmeta.NumToIDReq) (*blockmeta.BlockResp, error) {
	//todo: implement this method]
	log.Printf("Received following block num: %v", in.BlockNum)
	return &blockmeta.BlockResp{Id: "1234", Num: 1234, Time: timestamppb.New(time.Now())}, nil
}

func (s *serverBlock) IDToNum(ctx context.Context, in *blockmeta.IDToNumReq) (*blockmeta.BlockResp, error) {
	//todo: implement this method
	log.Printf("Received following block ID: %v", in.BlockID)
	return &blockmeta.BlockResp{Id: "1234", Num: 1234, Time: timestamppb.New(time.Now())}, nil
}

func (s *serverBlock) At(ctx context.Context, in *blockmeta.TimeReq) (*blockmeta.BlockResp, error) {
	//todo: implement this method
	log.Printf("Received following Timestamp: %v", in.Time)
	return &blockmeta.BlockResp{Id: "1234", Num: 1234, Time: timestamppb.New(time.Now())}, nil
}

func (s *serverBlock) After(ctx context.Context, in *blockmeta.RelativeTimeReq) (*blockmeta.BlockResp, error) {
	//todo: implement this method
	log.Printf("Received following Timestamp: %v", in.Time)
	return &blockmeta.BlockResp{Id: "1234", Num: 1234, Time: timestamppb.New(time.Now())}, nil
}

func (s *serverBlock) Before(ctx context.Context, in *blockmeta.RelativeTimeReq) (*blockmeta.BlockResp, error) {
	//todo: implement this method
	log.Printf("Received following Timestamp: %v", in.Time)
	return &blockmeta.BlockResp{Id: "1234", Num: 1234, Time: timestamppb.New(time.Now())}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sBlock := grpc.NewServer()
	blockmeta.RegisterBlockServer(sBlock, &serverBlock{})

	log.Printf("server block listening at %v", lis.Addr())
	if err := sBlock.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
