# User-RPC-Service

## Installation

Compile .proto file by folder_name.

```bash
protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/{folder_name}/*.proto
```

Run unit test by spesific unit.
```bash
Example:

go test -run ^TestUserAdd$ github.com/private-project-pp/product-rpc-service/usecase/users_administration
```

## Usage
```bash
make run
```