# api

## Compile Protobuf into Go

```
cd api
protoc -I . --go_out=plugins=grpc:. ./user.proto
```