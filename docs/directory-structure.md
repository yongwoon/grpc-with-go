# Project Structure

./src
├─ api
│ └─ hello.proto # proto ファイル
├─ cmd
│ └─ server
│ └─ main.go
├─ pkg
│ └─ grpc # 自動生成されたコード
│ ├─ hello.pb.go
│ └─ hello_grpc.pb.go
├─ go.mod
└─ go.sum

## 説明

- `src/api`
  - proto file が置かれる
- `src/cmd/server`
  - gRPC server file が置かれる
- `src/pkg/grpc`
  - `protoc` command で自動生成された file が置かれる
