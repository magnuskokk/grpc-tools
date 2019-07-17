.PHONY: all
all: generate

.PHONY: generate
generate: generate-grpc generate-http generate-client generate-doc
	$(MAKE) -C app generate

.PHONY: test
test:
	$(MAKE) -C app test

.PHONY: cleandoc
cleandoc:
	@rm -rf .direnv/swagger

.PHONY: clean
clean: cleandoc
	$(info - Removing all generated files and directories)
	$(MAKE) -C app clean
	$(MAKE) -C frontend clean
	$(MAKE) -C swagger clean

.PHONY: sudoclean
sudoclean: clean
	$(info - Force clean with .direnv removal)
	@sudo rm -rf ./.direnv

# TODO create a script go generate multiple services
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
./app \
		protobuf/services/heartbeat/service.proto

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
./app \
		protobuf/services/heartbeat/service.proto

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
		protobuf/services/heartbeat/service.proto

# TODO foreach multiple services
################################
# Generate swagger doc
################################
.PHONY: generate-doc
generate-doc:
	@rm -rf .direnv/swagger
	@mkdir -p .direnv/swagger

	$(info - Generate the swagger/pkg/server/service.swagger.json file)
	@protoc \
		-I. \
		-I${GOPATH}/src/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/gogo/protobuf/protobuf/ \
		--swagger_out=logtostderr=true:./swagger/ \
		protobuf/services/heartbeat/service.proto

	$(info - Install swagger-ui-dist static files)
	@npm install --prefix .direnv/ -g swagger-ui-dist

	$(info - Copy the swagger ui dist from node package dir)
	@cp -r ${GOPATH}/lib/node_modules/swagger-ui-dist/* .direnv/swagger/

	$(info - Copy services)
	@cp -r ./swagger/protobuf .direnv/swagger/protobuf

	$(info - Replace the default example with our own service documentation)
	@sed -i -e 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/http:\/\/localhost:8000\/docs\/protobuf\/services\/heartbeat\/service.swagger.json/g' .direnv/swagger/index.html
