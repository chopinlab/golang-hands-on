# grpc-hands-on for Golang

### GRPC
https://grpc.io/

```shell
brew install protobuf
protoc --version  # Ensure compiler version is 3+
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./pb/helloworld.proto
    
go get -u google.golang.org/grpc
``` 


