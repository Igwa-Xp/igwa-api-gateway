
# Generate proto files for gRPC server and client 
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/**/pb/*.proto

# Run gRPC server
server:
	go run cmd/main.go

	