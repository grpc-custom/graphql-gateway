red-server:
	go run ./cmd/main.go -port=9001 -svc=red

green-server:
	go run ./cmd/main.go -port=9002 -svc=green

gateway-server:
	go run ./main.go

protoc:
	protoc \
		-I=${GOPATH}/src:. \
		--go_out=plugin=grpc:. \
		--graphql-gateway_out=. \
		./proto/red/*.proto

.PHONY: red-server green-server gateway-server protoc
