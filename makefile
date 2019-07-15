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
# Generate documentation
################################
.PHONY: doc
doc:
	$(info Generating swagger doc)
	@mkdir -p ./swagger
	# Copy swagger-dist-ui
	@cp -r $GOPATH/lib/node_modules/swagger-ui-dist/* ./swagger/
	# Replace the default example with our own service.
	@sed -i -e 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/http:\/\/localhost:8000\/openapi-ui\/server.swagger.json/g' ./swagger/index.html
	# Generate the statik asset
	@statik -m -f -src ./swagger

.PHONY: all
all: clean protoc generate doc

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
	@go test -v ./...

################################
# Clean up generated files
################################
.PHONY: clean
clean:
	$(info Deleting all generated files and directories)
	@find ./pkg -type f -iname "*.pb.go" -exec rm {} \;
	@find ./pkg -type f -iname "*.pb.gw.go" -exec rm {} \;
	@rm -rf statik swagger
