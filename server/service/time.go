package service

import (
	pbkv "github.com/streamingfast/substreams-sink-kv/pb/substreams/sink/kv/v1"
)

type BlockByTime struct {
	sinkClient pbkv.KvClient
}

func NewBlockByTimeService(sinkClient pbkv.KvClient) *BlockByTime {
	return &BlockByTime{
		sinkClient: sinkClient,
	}
}
