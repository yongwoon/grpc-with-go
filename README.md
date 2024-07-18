# grpc-with-go

[作ってわかる！ はじめての gRPC](https://zenn.dev/hsaki/books/golang-grpc-starting) のコード練習

## Getting Start

- install protobuf

```bash
brew install protobuf
which protoc

# /usr/local/bin/protoc # コマンド配置箇所のパスが出力されればOK
```

- Setup package

```bash
cd src
go mod init grpc-with-go
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
```

- set GO path

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Create code with `protoc` command

- `--go_out`
  - hello.pb.go 파일의 출력 대상 디렉토리 지정
- `--go_opt`
  - hello.pb.go 파일 생성시 options.
  - 이번에는 paths=source_relative지정하여 --go_out옵션 지정이 상대 경로임을 명시
- `--go-grpc_out`
  - hello_grpc.pb.go파일의 출력 대상 디렉토리 지정
- `--go-grpc_opt`
  - hello_grpc.pb.go파일 생성시 옵션.
  - 이번에는 paths=source_relative 지정하여 --go-grpc_out옵션 지정이 상대 경로임을 명시

```bash
cd api
protoc --go_out=../pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../pkg/grpc --go-grpc_opt=paths=source_relative \
	hello.proto
```

## Utils

- install grpcurl

```bash
brew install grpcurl

which grpcurl
```

## Link

- [Language Guide (proto 3)](https://protobuf.dev/programming-guides/proto3/)
- [Protocol Buffers Well-Known Types](https://protobuf.dev/reference/protobuf/google.protobuf/)
