# Setup the dev environment.
.PHONY: setup
setup:
	@setup

# Generate gRPC server and client, HTTP REST API server, validators and swagger docs.
.PHONY: protoc
protoc:
	@protoc-service protobuf/heartbeat/service.proto

# Generate protos and mocks.
.PHONY: generate
generate: protoc
	@go generate -x ./...

# Generate static assets for OpenAPI UI.
.PHONY: statik
statik:
	@statik -m -f -src swagger/

test:
	@go test -v ./...

opendoc:
	@go run main.go
	@xdg-open localhost:8000/openapi-ui