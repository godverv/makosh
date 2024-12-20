gen: gen-server-grpc
build-local-container:
	docker buildx build \
			--load \
			--platform linux/arm64 \
			-t makosh:local .

### Grpc server generation
gen-server-grpc: .deps-grpc .gen-server-grpc

.deps-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/Red-Sock/protoc-gen-docs@latest
	EASYPPATH=proto_deps easyp mod download

.gen-server-grpc:
	EASYPPATH=proto_deps easyp generate