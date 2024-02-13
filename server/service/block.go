package service

import (
	pbkv "github.com/streamingfast/substreams-sink-kv/pb/substreams/sink/kv/v1"
)

type Block struct {
	sinkClient pbkv.KvClient
}

func NewBlockService(sinkClient pbkv.KvClient) *Block {
	return &Block{
		sinkClient: sinkClient,
	}
}
