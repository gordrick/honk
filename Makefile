include .env
export
proto:
	protoc --go_out=pkg/pb --go_opt=paths=source_relative \
            --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
            -I=pkg/pb \
            pkg/pb/**/*.proto
server:
	go run cmd/*