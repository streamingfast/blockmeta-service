package server

import (
	"log"
	"math/big"
	"time"

	"github.com/streamingfast/dgrpc"
	kvv1 "github.com/streamingfast/substreams-sink-kv/pb/substreams/sink/kv/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConnectToSinkServer(host string) kvv1.KvClient {
	conn, err := dgrpc.NewInternalClient(host)
	if err != nil {
		log.Fatalf("did not connect to the sink server: %v", err)
	}

	return kvv1.NewKvClient(conn)
}

func valueToTimestamp(timestampAsBytes []byte) (pbTimestamp *timestamppb.Timestamp) {
	timestampAsUnix := time.UnixMilli(big.NewInt(0).SetBytes(timestampAsBytes).Int64())
	pbTimestamp = timestamppb.New(timestampAsUnix)
	return pbTimestamp
}

func valueToBlockNumber(blockNumberAsBytes []byte) (blockNumber uint64) {
	blockNumber = big.NewInt(0).SetBytes(blockNumberAsBytes).Uint64()
	return blockNumber
}
