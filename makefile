# Setup the development environment.
# Setup development environment.
.PHONY: setup
setup:
	@setup

# Generate gRPC server and client, HTTP REST API server, validators and swagger.
.PHONY: protoc
protoc:
	@protoc-service protobuf/heartbeat/service.proto

# Generate protos and mocks.
.PHONY: generate
generate: protoc
	@go generate -x ./...

# Generate static assets for OpenAPI UI.
.PHONY: statik
swagger:
	statik -m -f -src third_party/OpenAPI/

test:
	@go test -v ./...

