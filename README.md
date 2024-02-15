# BlockMeta Service 

This repository provides all the keys to run a block meta service that provides high-level data on blocks, for any blockchain supported by StreamingFast.
To make such a service, you will have to extract block meta data first, store them in a kv store and then serve it.

## Extracting and Storing data 

To extract block meta data, you will need `substreams_spkg_path`. In order to get this package path, please visit the [substreams](./substreams/README.md). 
Then you will need to run a substreams-sink-kv server, please refer to the [substreams-sink-kv](https://github.com/streamingfast/substreams-sink-kv) 
repository that explains how to do so. 

This server will then be able to extract block meta data from the substreams and store it in a key-value store.

## Serving data

To serve block meta data, first be sure you have a substreams-sink-kv server running. 
Then, you can run a block-meta server referring to the [block-meta documentation](./server/README.md).

