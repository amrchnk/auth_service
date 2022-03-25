create:
	protoc --proto_path=proto proto/*.proto --go_out=./
	protoc --proto_path=proto proto/*.proto --go-grpc_out=./

run:
	go run ./cmd/main.go