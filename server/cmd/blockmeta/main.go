package main

import (
	"blockmeta_server/kvtool"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "blockmeta_server/pb/blockmeta"
	pbkv "blockmeta_server/pb/github.com/streamingfast/substreams-sink-kv/pb"
	"github.com/streamingfast/dgrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port       = flag.Int("port", 50051, "The server port")
	clientMode = flag.Bool("clientMode", false, "Run client mode for testing")
)

type serverBlock struct {
	pb.UnimplementedBlockServer
	pb.UnimplementedBlockByTimeServer
}

func (s *serverBlock) NumToID(ctx context.Context, in *pb.NumToIDReq) (*pb.BlockResp, error) {

	prefix := kvtool.PackNumPrefixKey(in.BlockNum)

	pbkvClient := connectToSinkServer()

	response, err := pbkvClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block number: %v", in.BlockNum)
	}

	blockNum, blockID, err := kvtool.UnpackNumIDKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockPbTimestamp := kvtool.UnpackTimeValue(response.KeyValues[0].Value)
	return &pb.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *serverBlock) IDToNum(ctx context.Context, in *pb.IDToNumReq) (*pb.BlockResp, error) {

	prefix := kvtool.PackIDPrefixKey(in.BlockID)

	pbkvClient := connectToSinkServer()

	response, err := pbkvClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block id: %v", in.BlockID)
	}

	blockNum, blockID, err := kvtool.UnpackIDNumKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockPbTimestamp := kvtool.UnpackTimeValue(response.KeyValues[0].Value)
	return &pb.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *serverBlock) At(ctx context.Context, in *pb.TimeReq) (*pb.BlockResp, error) {
	prefix := kvtool.PackTimePrefixKey(in.Time.AsTime(), true)

	pbkvClient := connectToSinkServer()

	response, err := pbkvClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block timestamp: %v", in.Time)
	}

	blockPbTimestamp, blockID, err := kvtool.UnpackTimeIDKey(response.KeyValues[0].Key, true)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockNum := kvtool.UnpackBlockNumberValue(response.KeyValues[0].Value)
	return &pb.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *serverBlock) Before(ctx context.Context, in *pb.RelativeTimeReq) (*pb.BlockResp, error) {
	prefix := kvtool.PackTimePrefixKey(in.Time.AsTime(), false)

	pbkvClient := connectToSinkServer()

	response, err := pbkvClient.Scan(ctx, &pbkv.ScanRequest{Begin: prefix, Limit: 1})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block timestamp: %v", in.Time)
	}

	blockPbTimestamp, blockID, err := kvtool.UnpackTimeIDKey(response.KeyValues[0].Key, false)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockNum := kvtool.UnpackBlockNumberValue(response.KeyValues[0].Value)
	return &pb.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *serverBlock) After(ctx context.Context, in *pb.RelativeTimeReq) (*pb.BlockResp, error) {
	prefix := kvtool.PackTimePrefixKey(in.Time.AsTime(), true)

	pbkvClient := connectToSinkServer()

	response, err := pbkvClient.Scan(ctx, &pbkv.ScanRequest{Begin: prefix, Limit: 1})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block timestamp: %v", in.Time)
	}

	blockPbTimestamp, blockID, err := kvtool.UnpackTimeIDKey(response.KeyValues[0].Key, true)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockNum := kvtool.UnpackBlockNumberValue(response.KeyValues[0].Value)
	return &pb.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func main() {
	flag.Parse()
	if !*clientMode {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		sBlock := grpc.NewServer()
		pb.RegisterBlockServer(sBlock, &serverBlock{})
		pb.RegisterBlockByTimeServer(sBlock, &serverBlock{})

		log.Printf("server block listening at %v", lis.Addr())
		if err := sBlock.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	} else {
		testClient()
	}
}

func connectToSinkServer() pbkv.KvClient {
	conn, err := dgrpc.NewInternalClient("localhost:7878")
	if err != nil {
		log.Fatalf("did not connect to the sink server: %v", err)
	}

	return pbkv.NewKvClient(conn)
}

func testClient() {
	conn, err := dgrpc.NewInternalClient("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewBlockClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	blockResp, err := c.NumToID(ctx, &pb.NumToIDReq{BlockNum: 6})
	if err != nil {
		log.Fatalf("could not get block data  NumToID request: %v", err)
	}

	blockSecondResp, err := c.IDToNum(ctx, &pb.IDToNumReq{BlockID: "f37c632d361e0a93f08ba29b1a2c708d9caa3ee19d1ee8d2a02612bffe49f0a9"})
	if err != nil {
		log.Fatalf("could not get block data for IDToNum request: %v", err)
	}

	d := pb.NewBlockByTimeClient(conn)
	blockThirdResp, err := d.At(ctx, &pb.TimeReq{Time: timestamppb.New(time.UnixMilli(1438270239000))})
	if err != nil {
		log.Fatalf("could not get block data for At request: %v", err)
	}

	blockFourthResp, err := d.After(ctx, &pb.RelativeTimeReq{Time: timestamppb.New(time.UnixMilli(1438270239090)), Inclusive: true})
	if err != nil {
		log.Fatalf("could not get block data for After request: %v", err)
	}

	blockFifthResp, err := d.Before(ctx, &pb.RelativeTimeReq{Time: timestamppb.New(time.UnixMilli(1438270239090)), Inclusive: true})
	if err != nil {
		log.Fatalf("could not get block data for Before request: %v", err)
	}

	fmt.Printf("Blockresponse for NumToID request: %s\n", blockResp.String())

	fmt.Printf("Blockresponse for  IDToNum request: %s\n", blockSecondResp.String())

	fmt.Printf("Blockresponse for  At request: %s\n", blockThirdResp.String())

	fmt.Printf("Blockresponse for After request: %s\n", blockFourthResp.String())

	fmt.Printf("Blockresponse for Before request: %s\n", blockFifthResp.String())
}
