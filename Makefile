protofmt:
	find . -name "*.proto" | xargs clang-format -i

protoc:
	protoc -I. \
		--go_out=plugins=grpc:${GOPATH}/src \
		graphql.proto

build-graphql-gateway-debug:
	go build -o protoc-gen-graphql-gateway-debug ./cmd/protoc-gen-graphql-gateway/main.go

graphql-gateway-debug: build-graphql-gateway-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
        -I=${GOPATH}/src/github.com/grpc-custom:. \
        --plugin=./protoc-gen-graphql-gateway-debug \
       	--graphql-gateway-debug_out=logtostderr=true:. \
        test/basic/user.proto

protoc-test-gql: build-graphql-gateway-debug
	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/sample/proto/green/*.proto

	protoc \
		-I=${GOPATH}/src:. \
		-I=${GOPATH}/src/github.com/googleapis/googleapis:. \
		-I=${GOPATH}/src/github.com/grpc-custom:. \
		--plugin=./protoc-gen-graphql-gateway-debug \
		--go_out=plugins=grpc:. \
		--graphql-gateway-debug_out=logtostderr=true:. \
		example/sample/proto/red/*.proto

.PHONY: gazelle protofmt protoc
