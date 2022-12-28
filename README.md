# api
api to communicate with the backend

## Getting Started with gRPC
https://grpc.io/docs/languages/go/quickstart/

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

```shell
protoc --go-grpc_out=grpc --go_out=grpc grpc/image_service.proto
protoc --go-grpc_out=grpc --go_out=grpc grpc/category_service.proto
```

## update dependencies

```shell
go get -u && go mod tidy
```