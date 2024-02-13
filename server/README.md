
## Overview
The Block Meta Service API provides a set of gRPC services for querying blockchain block metadata. It enables users to retrieve block IDs using block numbers,
convert block IDs to block numbers, fetch the latest block information, and query blocks by specific timestamps.

## Endpoints

To interact with the Block Meta Service, clients should use the gRPC endpoints provided by streamingFast for all the chains

## Example Query and Response

### Querying block information using a block number

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -d '{"blockNum": "501"}' holesky.eth.streamingfast.io:443 sf.blockmeta.v2.Block/NumToID
```

```json
{
"id": "00cd01a162b0bc1e9a88eb8718891dd31984f3fb64c2392570d310d8a5f05bf6",
"num": "501",
"time": "2023-09-28T14:16:24Z"
}
```

### Querying block information using a block ID

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -d '{"blockID": "0000000046887292a76cd113a5fd6af38b17c9fb77e5936cd9856694030598f9"}' mainnet.btc.streamingfast.io:443 sf.blockmeta.v2.Block/IDToNum
```

```json
{
"id": "0000000046887292a76cd113a5fd6af38b17c9fb77e5936cd9856694030598f9",
"num": "501",
"time": "2009-01-14T21:38:31Z"
}
```

### Querying head block information

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" .eth.streamingfast.io:443 sf.blockmeta.v2.Block/Head
```

```json
{
"id": "50446fbfca349144a5346106038c1865e753cf138db9d3f6ef25224f75c198e9",
"num": "931897",
"time": "2024-02-12T22:12:24Z"
}
```

### Querying block information at a specific timestamp

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -d '{"time": "2024-02-12T19:17:23Z"}' mainnet.eth.streamingfast.io:443 sf.blockmeta.v2.BlockByTime/At

```

```json
{
"id": "1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748",
"num": "19213975",
"time": "2024-02-12T19:17:23Z"
}
```

### Querying block information after a specific timestamp by setting inclusive to true

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -d '{"time": "2024-02-12T19:17:23Z", "inclusive": "true"}' mainnet.eth.streamingfast.io:443 sf.blockmeta.v2.BlockByTime/After
```

```json
{
"id": "1cddf333f43b88edb8dcf30861542f13297e3e5a90fd03ec044926c3440ea748",
"num": "19213975",
"time": "2024-02-12T19:17:23Z"
}
```
### Querying block information before a specific timestamp by setting inclusive to false

```
grpcurl -H "Authorization: Bearer $YOUR_TOKEN_HERE" -d '{"time": "2024-02-12T19:17:23Z", "inclusive": "false"}' mainnet.eth.streamingfast.io:443 sf.blockmeta.v2.BlockByTime/Before```
```

```json
{
"id": "4132f03a2c4bf07be79b0d99e7d21aa2cdf71486b415e6031c0b3a28fd33fe2a",
"num": "19213974",
"time": "2024-02-12T19:17:11Z"
}
```
