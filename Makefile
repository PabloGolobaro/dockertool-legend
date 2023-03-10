server:
	go run ./cmd/server
client:
	go run ./cmd/client -timeout=15
grpc:
	protoc --proto_path=./api --go_out=./internal/api --go-grpc_out=./internal/api ./api/*.proto
build:
	go build -o ./bin/legtool.exe ./cmd/server
	go build -o ./bin/legtoolclient.exe ./cmd/client

server_debug:
	go run -tags dev ./cmd/server
