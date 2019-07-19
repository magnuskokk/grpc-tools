.PHONY: all
all: generate

.PHONY: generate
generate:
	generate-grpc ${REPO_ROOT}/protobuf/services/echo/service.proto
	generate-grpc ${REPO_ROOT}/protobuf/services/ping/service.proto

	generate-http ${REPO_ROOT}/protobuf/services/echo/service.proto
	generate-http ${REPO_ROOT}/protobuf/services/ping/service.proto

	generate-client ${REPO_ROOT}/protobuf/services/echo/service.proto
	generate-client ${REPO_ROOT}/protobuf/services/ping/service.proto

	generate-swagger ${REPO_ROOT}/protobuf/services/echo/service.proto
	generate-swagger ${REPO_ROOT}/protobuf/services/ping/service.proto

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

