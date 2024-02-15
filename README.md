## BlockMeta Service 

This repository provides all the keys to run a block meta service that provides high-level data on blocks, for any blockchain supported by StreamingFast.
To make such a service, extract block meta data first, store them in a kv store and then serve it.

### Extracting and Storing data 

- Run a [substreams-sink-kv](https://github.com/streamingfast/substreams-sink-kv) server, providing a [substreams_spkg_path](./substreams/README.md).

### Serving data

- Run a [block-meta server](./server/README.md), connecting this server to the *substreams-sink-kv* running server.

