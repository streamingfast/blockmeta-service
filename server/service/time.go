package service

import (
	"context"
	"fmt"
	"log/slog"

	pbbmsrv "github.com/streamingfast/blockmeta-service/pb/sf/blockmeta/v2"
	pbkv "github.com/streamingfast/blockmeta-service/pb/sf/substreams/sink/kv/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BlockByTime struct {
	sinkClient pbkv.KvClient
}

func NewBlockByTimeService(sinkClient pbkv.KvClient) *BlockByTime {
	return &BlockByTime{
		sinkClient: sinkClient,
	}
}

func (s *BlockByTime) At(ctx context.Context, in *pbbmsrv.TimeReq) (*pbbmsrv.BlockResp, error) {
	slog.Info("handling At request", "block_time", in.Time)
	prefix := Keyer.PackTimePrefixKey(in.Time.AsTime(), false)

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block timestamp: %v", in.Time)
	}

	blockPbTimestamp, blockID, err := Keyer.UnpackTimeIDKey(response.KeyValues[0].Key, false)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockNum := valueToBlockNumber(response.KeyValues[0].Value)
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *BlockByTime) Before(ctx context.Context, in *pbbmsrv.RelativeTimeReq) (*pbbmsrv.BlockResp, error) {
	slog.Info("handling Before request", "block_time", in.Time)
	prefix := Keyer.PackTimePrefixKey(in.Time.AsTime(), false)

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
			return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
		}

		if !in.Inclusive && (blockPbTimestamp.AsTime() == in.Time.AsTime()) {
			continue
		}

		blockNum = valueToBlockNumber(response.KeyValues[i].Value)
		break
	}
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *BlockByTime) After(ctx context.Context, in *pbbmsrv.RelativeTimeReq) (*pbbmsrv.BlockResp, error) {
	slog.Info("handling After request", "block_time", in.Time)
	prefix := Keyer.PackTimePrefixKey(in.Time.AsTime(), true)

	response, err := s.sinkClient.Scan(ctx, &pbkv.ScanRequest{Begin: prefix, Limit: 4})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	var blockID string
	var blockNum uint64
	blockPbTimestamp := &timestamppb.Timestamp{}

	for i := 0; i < len(response.KeyValues); i++ {

		blockPbTimestamp, blockID, err = Keyer.UnpackTimeIDKey(response.KeyValues[i].Key, true)
		if err != nil {
			return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
		}

		if !in.Inclusive && (blockPbTimestamp.AsTime() == in.Time.AsTime()) {
			continue
		}

		blockNum = valueToBlockNumber(response.KeyValues[i].Value)

		break
	}
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}
