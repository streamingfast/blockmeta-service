package service

import (
	"log"
	"math/big"
	"time"

	pbkv "github.com/streamingfast/blockmeta-service/pb/sf/substreams/sink/kv/v1"
	"github.com/streamingfast/dgrpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConnectToSinkServer(host string) pbkv.KvClient {
	conn, err := dgrpc.NewInternalClient(host)
	if err != nil {
		log.Fatalf("did not connect to the sink server: %v", err)
	}

	return pbkv.NewKvClient(conn)
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
