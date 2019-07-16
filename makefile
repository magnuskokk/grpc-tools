.PHONY: all
all: generate

.PHONY: generate
generate: generate-grpc generate-http generate-client generate-doc
	$(MAKE) -C backend generate

.PHONY: test
test:
	$(MAKE) -C backend test

.PHONY: clean
clean:
	$(info - Removing all generated files and directories)
	$(MAKE) -C backend clean
	$(MAKE) -C frontend clean
	@rm -rf swagger

.PHONY: sudoclean
sudoclean: clean
	$(info - Force clean with .direnv removal)
	@sudo rm -rf ./.direnv

.PHONY: docserver
docserver:
	$(info - Serving documentation at: http://localhost:8000/openapi-ui)
	@go run cmd/docserver/main.go

################################
# Generate gRPC server and client
################################
.PHONY: generate-grpc
generate-grpc:
	$(info - Generate grpc server and client)
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
./backend \
		backend/pkg/server/service.proto

################################
# Generate HTTP gateway for gRPC server
################################
.PHONY: generate-http
generate-http:
	$(info - Generating http gateway)
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
./backend \
		backend/pkg/server/service.proto

################################
# Generate frontend client for gRPC server
# TODO: fix paths and typescript imports
################################
.PHONY: generate-client
generate-client:
	$(info - Generate typescript client)
	@mkdir -p ./frontend/generated
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:\
./frontend/generated \
		backend/pkg/server/service.proto

################################
# Generate swagger doc
################################
.PHONY: generate-doc
generate-doc:
	@rm -rf ./swagger
	@mkdir -p ./swagger

	$(info - Generate the swagger/pkg/server/service.swagger.json file)
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--swagger_out=logtostderr=true:swagger/ \
		backend/pkg/server/service.proto

	$(info - Install swagger-ui-dist static files)
	@npm install --prefix ${GOPATH} -g swagger-ui-dist

	$(info - Copy the swagger ui dist from node package dir)
	@cp -r ${GOPATH}/lib/node_modules/swagger-ui-dist/* ./swagger/

	$(info - Replace the default example with our own service documentation)
	@sed -i -e 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/http:\/\/localhost:8000\/openapi-ui\/pkg\/server\/service.swagger.json/g' ./swagger/index.html

	$(info - Pack the doc web application into single file)
	@statik -m -f -src ./swagger -dest ./backend
