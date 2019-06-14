# api

## Compile Protobuf into Go

```
cd api
protoc -I . --go_out=plugins=grpc:. ./user.proto
protoc -I . --go_out=plugins=grpc:. ./role.proto
```

### Add Custom Tags to Protobuf

https://github.com/srikrsna/protoc-gen-gotag
go get github.com/srikrsna/protoc-gen-gotag

```
protoc -I %GOPATH%\src -I . --gotag_out=:. ./role.proto
```
