## Generate .proto file.

Download the precompiled binary from https://github.com/google/protobuf/releases
Copy the protoc binary inside the **myservice** directory

```
./protoc --go_out=plugins=grpc:. myservice.proto
