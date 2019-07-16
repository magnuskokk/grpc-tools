.PHONY: all
all: protoc generate doc

################################
# Run protoc with preconfigured plugins
################################
.PHONY: protoc
protoc:
	$(info Generating go code for gRPC server, client, http gateway and swagger docs)
	@protoc-service pkg/server/service.proto

################################
# Generate mocks
################################
.PHONY: generate
generate:
	$(info Running //go:generate directives) 
	@go generate -x ./...

################################
# Generate and overwrite with fresh generated documentation
################################
.PHONY: doc
doc:
	$(info Generating swagger doc for pkg/server)
	@swagger

################################
# Serve docs at http://localhost:8000/openapi-ui
################################
.PHONY: docserver
docserver:
	$(info Serving gRPC API documentation at: http://localhost:8000/openapi-ui)
	@go run cmd/docserver/main.go
	
################################
# Run all tests
################################
.PHONY: test
test:
	$(info Running all tests)
	@go test -count=1 -v ./...

################################
# Run all benchmarks
################################
.PHONY: bench
bench:
	$(info Running all benchmarks)
	@go test -run=xxx -bench=. ./...

################################
# Clean up generated files
################################
.PHONY: clean
clean:
	$(info Deleting all generated files and directories)
	@find ./pkg -type f -iname "*.pb.go" -delete
	@find ./pkg -type f -iname "*.pb.gw.go" -delete
	@find ./pkg -type d -iname "mocks" -exec rm -rf {} +
	@rm -rf statik swagger

.PHONY: sudoclean
sudoclean: clean
	$(info Force clean with .direnv removal. Must have GOPATH set to ./.direnv and setup run before any make commands)
	@sudo rm -rf ./.direnv
