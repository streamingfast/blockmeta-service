package kvtool

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"math/big"
	"strings"
	"time"
)

const (
	TblPrefixBlockIDs    = "1"
	TblPrefixBlockNums   = "2"
	TblPrefixTimelineFwd = "3"
	TblPrefixTimelineBck = "4"
)

func PackNumPrefixKey(blockNum uint64) string {
	keyPrefix := TblPrefixBlockNums + ":" + packU64Reverse(blockNum) + ":"
	return keyPrefix
}

func PackIDPrefixKey(id string) string {
	keyPrefix := TblPrefixBlockIDs + ":" + id + ":"
	return keyPrefix
}

func PackTimePrefixKey(time time.Time, fwd bool) string {
	timeAsUnixMillis := uint64(time.UnixMilli())
	if fwd {
		return TblPrefixTimelineFwd + ":" + packU64(timeAsUnixMillis) + ":"
	}
	return TblPrefixTimelineBck + ":" + packU64Reverse(timeAsUnixMillis) + ":"
}

func packU64Reverse(num uint64) string {
	return packU64(math.MaxUint64 - num)
}

func packU64(num uint64) string {
	numBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(numBytes, num)
	return hex.EncodeToString(numBytes)
}

func UnpackNumIDKey(key string) (blockNum uint64, blockID string, err error) {
	parts := strings.Split(key, ":")

	parts = parts[1:] // remove the prefix

	if len(parts) != 2 {
		err = fmt.Errorf("invalid key: %s", key)
		return 0, "", err
	}

	encodedReversedBlockNum := parts[0]
	decodedBlockNumBytes, err := hex.DecodeString(encodedReversedBlockNum)
	if err != nil {
		return 0, "", fmt.Errorf("error decoding block number: %w", err)
	}
	reverseBlockNum := binary.BigEndian.Uint64(decodedBlockNumBytes)

	blockID = parts[1]
	blockNum = math.MaxUint64 - reverseBlockNum

	return blockNum, blockID, nil
}

func UnpackIDNumKey(key string) (blockNum uint64, blockID string, err error) {
	parts := strings.Split(key, ":")

	parts = parts[1:] // remove the prefix

	if len(parts) != 2 {
		err = fmt.Errorf("invalid key: %s", key)
		return 0, "", err
	}

	encodedReversedBlockNum := parts[1]
	decodedBlockNumBytes, err := hex.DecodeString(encodedReversedBlockNum)
	if err != nil {
		return 0, "", fmt.Errorf("error decoding block number: %w", err)
	}
	reverseBlockNum := binary.BigEndian.Uint64(decodedBlockNumBytes)

	blockID = parts[0]
	blockNum = math.MaxUint64 - reverseBlockNum

	return blockNum, blockID, nil
}

func UnpackTimeIDKey(key string, fwd bool) (pbTimestamp *timestamppb.Timestamp, blockID string, err error) {
	parts := strings.Split(key, ":")

	parts = parts[1:] // remove the prefix

	if len(parts) != 2 {
		err = fmt.Errorf("invalid key: %s", key)
		return pbTimestamp, blockID, err
	}

	encodedTime := parts[0]
	decodedTimeBytes, err := hex.DecodeString(encodedTime)
	if err != nil {
		err = fmt.Errorf("error decoding time: %w", err)
		return pbTimestamp, blockID, err
	}
	decodedTime := binary.BigEndian.Uint64(decodedTimeBytes)

	blockID = parts[1]

	if !fwd {
		reverseTime := int64(math.MaxUint64 - decodedTime)
		pbTimestamp = timestamppb.New(time.UnixMilli(reverseTime))
		return pbTimestamp, blockID, nil
	}

	pbTimestamp = timestamppb.New(time.UnixMilli(int64(decodedTime)))

	return pbTimestamp, blockID, nil
}

func UnpackTimeValue(timestampAsBytes []byte) (pbTimestamp *timestamppb.Timestamp) {
	timestampAsUnix := time.UnixMilli(big.NewInt(0).SetBytes(timestampAsBytes).Int64())
	pbTimestamp = timestamppb.New(timestampAsUnix)
	return pbTimestamp
}

func UnpackBlockNumberValue(blockNumberAsBytes []byte) (blockNumber uint64) {
	blockNumber = big.NewInt(0).SetBytes(blockNumberAsBytes).Uint64()
	return blockNumber
}
