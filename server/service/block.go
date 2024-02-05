package service

import (
	"context"
	"fmt"

	pbbmsrv "github.com/streamingfast/blockmeta-service/pb/sf/blockmeta/v2"
	pbkv "github.com/streamingfast/blockmeta-service/pb/sf/substreams/sink/kv/v1"
)

type Block struct {
	sinkClient pbkv.KvClient
}

func NewBlockService(sinkClient pbkv.KvClient) *Block {
	return &Block{
		sinkClient: sinkClient,
	}
}

func (s *Block) NumToID(ctx context.Context, in *pbbmsrv.NumToIDReq) (*pbbmsrv.BlockResp, error) {
	prefix := Keyer.PackNumPrefixKey(in.BlockNum)

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block number: %v", in.BlockNum)
	}

	blockNum, blockID, err := Keyer.UnpackNumIDKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockPbTimestamp := valueToTimestamp(response.KeyValues[0].Value)
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *Block) IDToNum(ctx context.Context, in *pbbmsrv.IDToNumReq) (*pbbmsrv.BlockResp, error) {

	prefix := Keyer.PackIDPrefixKey(in.BlockID)

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	if len(response.KeyValues) > 1 {
		return nil, fmt.Errorf("more than one block found for block id: %v", in.BlockID)
	}

	blockNum, blockID, err := Keyer.UnpackIDNumKey(response.KeyValues[0].Key)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockPbTimestamp := valueToTimestamp(response.KeyValues[0].Value)
	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}

func (s *Block) Head(ctx context.Context, in *pbbmsrv.Empty) (*pbbmsrv.BlockResp, error) {
	prefix := TblPrefixTimelineBck + ":"

	response, err := s.sinkClient.GetByPrefix(ctx, &pbkv.GetByPrefixRequest{Prefix: prefix, Limit: 1})
	if err != nil {
		return nil, fmt.Errorf("error getting block data from sink server: %w", err)
	}

	blockPbTimestamp, blockID, err := Keyer.UnpackTimeIDKey(response.KeyValues[0].Key, false)
	if err != nil {
		return nil, fmt.Errorf("error unpacking block number and block ID: %w", err)
	}

	blockNum := valueToBlockNumber(response.KeyValues[0].Value)

	return &pbbmsrv.BlockResp{Id: blockID, Num: blockNum, Time: blockPbTimestamp}, nil
}
