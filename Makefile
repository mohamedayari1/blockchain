.PHONY: build run test proto

build:
	go build -o bin/blocker
run: build 
	./bin/docker
test:
	go test -v ./...



proto:
	protoc --proto_path=proto \
		--go_out=proto --go_opt=paths=source_relative \
		--go-grpc_out=proto --go-grpc_opt=paths=source_relative \
		proto/*.proto