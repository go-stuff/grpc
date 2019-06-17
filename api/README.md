# api

## Install protoc

This will compile a protobuf(`.proto`) file into a go (`.pb.go`) file.
- [github.com/golang/protobuf](https://github.com/golang/protobuf)

```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Install protoc-go-inject-tag

This will inject custom tags like `bson` or `json` into the already generated go files.
- [github.com/favadi/protoc-go-inject-tag](https://github.com/favadi/protoc-go-inject-tag)

```
go get github.com/favadi/protoc-go-inject-tag
```

## Example of how to Compile Protobuf into Go with Custom Tags

This examples shows how to compile the `proto` file then how to inject the `custom tags`. First `cd` into the directory that contains the `.proto` files, the first command will generate a `.pb.go` file the second will inject tags from comments in the `.proto` file with the following format `// @inject_tag: bson:"_id"`.

```
cd api

protoc -I . --go_out=plugins=grpc:. ./session.proto
protoc-go-inject-tag -input=./session.pb.go

protoc -I . --go_out=plugins=grpc:. ./permission.proto
protoc-go-inject-tag -input=./permission.pb.go

protoc -I . --go_out=plugins=grpc:. ./user.proto
protoc-go-inject-tag -input=./user.pb.go

protoc -I . --go_out=plugins=grpc:. ./role.proto
protoc-go-inject-tag -input=./role.pb.go

protoc -I . --go_out=plugins=grpc:. ./route.proto
protoc-go-inject-tag -input="./route.pb.go"
```
