.PHONY: all
all: generate-grpc generate-http generate-doc go-generate 

################################
# Generate gRPC server and client
################################
.PHONY: generate-grpc
generate-grpc:
	$(info Generate pkg/server/service.pb.go)
	$(info Generate pkg/server/servicepb_test.go)
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--gogoslick_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
./ \
		pkg/server/service.proto

################################
# Generate HTTP gateway for gRPC server
################################
.PHONY: generate-http
generate-http:
	$(info Generate pkg/server/service.pb.gw.go)
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--grpc-gateway_out=allow_patch_feature=false,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
./ \
		pkg/server/service.proto

################################
# Generate swagger doc
################################
.PHONY: generate-doc
generate-doc:
	@rm -rf ./swagger
	@mkdir -p ./swagger

	$(info Generate the swagger/pkg/server/service.swagger.json file)
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--swagger_out=logtostderr=true:swagger/ \
		pkg/server/service.proto

	$(info Install swagger-ui-dist static files)
	@npm install --prefix ${GOPATH} -g swagger-ui-dist

	$(info Copy the swagger ui dist from node package dir)
	@cp -r ${GOPATH}/lib/node_modules/swagger-ui-dist/* ./swagger/

	$(info Replace the default example with our own service documentation)
	@sed -i -e 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/http:\/\/localhost:8000\/openapi-ui\/pkg\/server\/service.swagger.json/g' ./swagger/index.html

	$(info Pack the doc web application into single file)
	@statik -m -f -src ./swagger

################################
# Generate code from //go:generate directives
################################
.PHONY: go-generate
go-generate:
	$(info Running //go:generate directives) 
	@go generate -x ./...

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
	@find ./pkg -type f -iname "*pb_test.go" -delete
	@find ./pkg -type d -iname "mocks" -exec rm -rf {} +
	@rm -rf statik swagger

.PHONY: sudoclean
sudoclean: clean
	$(info Force clean with .direnv removal. Must have GOPATH set to ./.direnv and setup run before any make commands)
	@sudo rm -rf ./.direnv
