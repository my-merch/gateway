generate:
	protoc -I ./proto \
      --go_out=paths=import:./stubs \
      --go_opt=paths=source_relative \
      --go-grpc_out=paths=import:./stubs \
      --go-grpc_opt=paths=source_relative \
      --grpc-gateway_out=paths=import:./stubs \
      --grpc-gateway_opt=paths=source_relative \
      ./proto/user-service/v1/users.proto
