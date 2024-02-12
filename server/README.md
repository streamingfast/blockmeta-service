# Block Meta Service API Documentation

## Overview

The Block Meta Service API provides a set of gRPC services for querying blockchain block metadata. It enables users to retrieve block IDs using block numbers, convert block IDs to block numbers, fetch the latest block information, and query blocks by specific timestamps.

## Services

### Block Service

The `Block` service offers functionalities to map block numbers to block IDs, block IDs to their corresponding block numbers, and to retrieve the latest block information.

#### RPC Methods

- `NumToID(NumToIDReq)`: Given a block number, returns the block ID, number, and timestamp.
- `IDToNum(IDToNumReq)`: Given a block ID, returns the block ID, number, and timestamp.
- `Head(Empty)`: Returns the latest block's ID, number, and timestamp.

### BlockByTime Service

The `BlockByTime` service provides capabilities to query block information based on specific timestamps.

#### RPC Methods

- `At(TimeReq)`: Returns the block closest to a specified timestamp.
- `After(RelativeTimeReq)`: Returns the first block after a specified timestamp (or the block at the specified timestamp if it exists, if the query is inclusive).
- `Before(RelativeTimeReq)`: Returns the last block before a specified timestamp (or the block at the specified timestamp if it exists, if the query is inclusive).

## Message Types

### Requests

- `NumToIDReq`: Request containing a block number (field: `blockNum`, type: `uint64`).
- `IDToNumReq`: Request containing a block ID (field: `blockID`, type: `string`).
- `TimeReq`: Request containing a specific timestamp (field: `time`, type: `google.protobuf.Timestamp`).
- `RelativeTimeReq`: Request containing a timestamp ((field: `time` , type: `google.protobuf.Timestamp`) , (field:`inclusive`, type: `bool`)) and a boolean indicating whether the query is inclusive of the given timestamp.

### Responses

- `BlockResp`: Response containing a block's ID (`string`), number (`uint64`), and timestamp (`google.protobuf.Timestamp`).

## Endpoints

To interact with the Block Meta Service, clients should use the following gRPC endpoints:

- `BlockServiceEndpoint`:
- `BlockByTimeServiceEndpoint`:

## Example Query and Response

### Querying block information using a block number

```
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto   -d '{"blockNum": "501"}' localhost:9000 sf.blockmeta.v2.Block/NumToID
```

```json
{
}
```

### Querying block information using a block ID

```
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto  -d '{"blockID": "501"}' localhost:9000 sf.blockmeta.v2.Block/IDToNum 
```

```json
{
}
```

### Querying head block information

```
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto localhost:9000 sf.blockmeta.v2.Block/Head
```

```json
{
}
```

### Querying block information at a specific timestamp

```  
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto -d '{"time": ""}' localhost:9000 sf.blockmeta.v2.BlockByTime/At

```

```json
{
}
```

### Querying block information after a specific timestamp by setting inclusive to true

```
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto -d '{"time": "", "inclusive": "true"}' localhost:9000 sf.blockmeta.v2.BlockByTime/After
```

```json
{
}
```
### Querying block information before a specific timestamp by setting inclusive to false

```
grpcurl --plaintext -proto server/proto/sf/blockmeta/v2/blockmeta.proto -d '{"time": , "inclusive": "false"}' localhost:9000 sf.blockmeta.v2.BlockByTime/Before
```

```json
{
}
```


//TODO: DO NOT FORGET TO SAY THIS API IS CHAIN AGNOSTIC 