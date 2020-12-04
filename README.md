# Nitro Examples

This is a repository for Nitro examples. Feel free to contribute.

## Contents

- [broker](broker) - A example of using Broker for Publish and Subscribing.
- [client](client) - Usage of the Client package to call a service.
- [config](config) - Using Go Config for dynamic config
- [graceful](graceful) - Demonstrates graceful shutdown of a service
- [helloworld](helloworld) - Hello world using micro
- [noproto](noproto) - Use micro without protobuf or code generation, only go types
- [options](options) - Setting options in the go-micro framework
- [server](server) - Use of the Server package directly to server requests.
- [service](service) - Example of the top level Service in go-micro.
- [sharding](sharding) - An example of how to shard requests or use session affinity
- [shutdown](shutdown) - Demonstrates graceful shutdown via context cancellation
- [stream](stream) - An example of a streaming service and client
- [waitgroup](waitgroup) - Demonstrates how to use a waitgroup with a service

## Install

Install [protoc](https://github.com/google/protobuf) for your environment. Then:

```shell
# install protoc-gen-go
go get github.com/golang/protobuf/{proto,protoc-gen-go}
# install protoc-gen-micro
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
```

To recompile any proto after changes:

```shell
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. path/to/proto
```
