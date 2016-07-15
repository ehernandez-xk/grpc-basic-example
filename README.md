# grpc-basic-example

This is a very small example using [grpc](http://www.grpc.io/) the client and server is written in [Go](https://golang.org/)

## Install

## Test it

```
go run server/main.go
go run client/main.go
```

## Generate myservice.pb.go file.

Download the precompiled binary from [github.com/google/protobuf](https://github.com/google/protobuf/releases) releases

Copy the protoc binary inside the **myservice** directory

```
./protoc --go_out=plugins=grpc:. myservice.proto
```
