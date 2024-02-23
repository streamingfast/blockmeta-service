# Table of Contents
- [Overview](#overview)
- [Running](#running)
- [Examples Queries and Responses](#examples-queries-and-responses)
    - [Block Service](#block-service)
        - [Querying block information using a block number](#querying-block-information-using-a-block-number)
        - [Querying block information using a block ID](#querying-block-information-using-a-block-id)
        - [Querying head block information](#querying-head-block-information)
    - [BlockByTime Service](#blockbytime-service)
        - [Querying block information at a specific timestamp](#querying-block-information-at-a-specific-timestamp)
        - [Querying block information after a specific timestamp by setting inclusive to true](#querying-block-information-after-a-specific-timestamp-by-setting-inclusive-to-true)
        - [Querying block information before a specific timestamp by setting inclusive to false](#querying-block-information-before-a-specific-timestamp-by-setting-inclusive-to-false)

## Overview
The Block Meta Service API provides a set of gRPC services for querying blockchain block metadata. It enables users to retrieve block IDs using block numbers,
convert block IDs to block numbers, fetch the latest block information, and query blocks by specific timestamps.

## Running 

Running a block-meta server, uses the `blockmeta` binary. 
In order to run properly, the server has to be connected to a [substreams-sink-kv](https://github.com/streamingfast/substreams-sink-kv) server.
To connect to the sink server, you can provide the substreams-sink-kv server address within the `--sink-addr` flag. 
The server will then connect to the sink server and start listening for gRPC requests on the address provided in the `--grpc-listen-addr` flag.

```bash 
blockmeta --sink-addr  localhost:9000 --grpc-listen-addr localhost:50051 
```

## Examples Queries and Responses

### Block Service 

#### Querying block information using a block number

```bash
curl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -H "Content-Type: application/json" --data '{"blockNum": "2"}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.Block/NumToID
```

```json
{
"id":"b495a1d7e6663152ae92708da4843337b958146015a2802f4193a410044698c9",
"num":"2",
"time":"2015-07-30T15:26:57Z"
}
```

#### Querying block information using a block ID

```bash
curl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -H "Content-Type: application/json" --data '{"blockID": "1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748"}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.Block/IDToNum 
```

```json
{
"id":"1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748",
"num":"19213975",
"time":"2024-02-12T19:17:23Z"}
```

#### Querying head block information

```bash
curl -H "Authorization: Bearer $SUBSTREAMS_API_TOKEN" -H "Content-Type: application/json" --data '{}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.Block/Head
```

```json
{
"id":"23b018bf48ee007187d776e4b1095d5a8c07e7db7bda60d5c9751fa666b8de84",
"num":"19293009",
"time":"2024-02-23T21:35:47Z"
}
```

### BlockByTime Service

#### Querying block information at a specific timestamp

```bash
curl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -H "Content-Type: application/json" --data '{"time": "2024-02-12T19:17:23Z"}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.BlockByTime/At
```

```json
{
"id": "1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748",
"num": "19213975",
"time": "2024-02-12T19:17:23Z"
}
```

#### Querying block information after a specific timestamp by setting inclusive to true

```bash
curl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -H "Content-Type: application/json" --data '{"time": "2024-02-12T19:17:23Z", "inclusive": true}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.BlockByTime/After
```

```json
{
"id":"1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748",
"num":"19213975",
"time":"2024-02-12T19:17:23Z"
}
```

#### Querying block information before a specific timestamp by setting inclusive to false

```bash
curl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -H "Content-Type: application/json" --data '{"time": "2024-02-12T19:17:23Z", "inclusive": false}' https://mainnet.eth.streamingfast.io:443/sf.blockmeta.v2.BlockByTime/Before
```

```json
{
"id":"4132f03a2c4bf07be79b0d99e7d21aa2cdf71486b415e6031c0b3a28fd33fe2a",
"num":"19213974",
"time":"2024-02-12T19:17:11Z"
}
```
