package server

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	KeyPrefixBlockTimeByBlockID     = "1"
	KeyPrefixBlockTimeByBlockNumber = "2"
	KeyPrefixBlockNumberByTimeFwd   = "3"
	KeyPrefixBlockNumberByTimeBwd   = "4"
)

var Keyer keyer

type keyer struct{}

func (keyer) PackNumPrefixKey(blockNum uint64) string {
	keyPrefix := KeyPrefixBlockTimeByBlockNumber + ":" + packU64Reverse(blockNum)
	return keyPrefix
}

func (keyer) PackBlockTimeByBlockIDKeyPrefix(id string) string {
	keyPrefix := KeyPrefixBlockTimeByBlockID + ":" + id + ":"
	return keyPrefix
}

func (keyer) PackTimePrefixKey(time time.Time, fwd bool) string {
	timeAsUnixMillis := uint64(time.UnixMilli())
	if fwd {
		return KeyPrefixBlockNumberByTimeFwd + ":" + packU64(timeAsUnixMillis)
	}
	return KeyPrefixBlockNumberByTimeBwd + ":" + packU64Reverse(timeAsUnixMillis)
}

func (keyer) UnpackNumIDKey(key string) (blockNum uint64, blockID string, err error) {
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

func (keyer) UnpackIDNumKey(key string) (blockNum uint64, blockID string, err error) {
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

func (keyer) UnpackTimeIDKey(key string, fwd bool) (pbTimestamp *timestamppb.Timestamp, blockID string, err error) {
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

func packU64Reverse(num uint64) string {
	return packU64(math.MaxUint64 - num)
}

func packU64(num uint64) string {
	numBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(numBytes, num)
	return hex.EncodeToString(numBytes)
}
