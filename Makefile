create:
	protoc --proto_path=proto proto/*.proto --go_out=./
	protoc --proto_path=proto proto/*.proto --go-grpc_out=./

run:
	go run ./cmd/main.go

goose up:
	goose postgres "user=postgres password=postgres dbname=auth_service sslmode=disable" up

goose down:
	goose postgres "user=postgres password=postgres dbname=auth_service sslmode=disable" down

docker:
	docker-compose up