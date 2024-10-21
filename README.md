# HyperFS

Scalable Infinite Storage Server.

## Binaries

-   [cmd/hyperfs](cmd/hyperfs): The `hyperfs` command line
-   [cmd/hyperfs-api](cmd/hyperfs-api): The `hyperfs-api` is a http api server
-   [cmd/hyperfs-index](cmd/hyperfs-index): The `hyperfs-index` is a metadata server
-   [cmd/hyperfs-storage](cmd/hyperfs-storage): The `hyperfs-storage` is a node of storage server

## Libraries

-   [pkg/hyperfs](pkg/hyperfs): The `hyperfs` is a core library

## Installation

```bash
#For Apple silicon
wget https://github.com/apple/foundationdb/releases/download/7.3.52/FoundationDB-7.3.52_arm64.pkg
#For Apple intel
wget https://github.com/apple/foundationdb/releases/download/7.3.52/FoundationDB-7.3.52_x86_64.pkg
```

## License

HyperFS is licensed under [the MIT license](LICENSE.md).
