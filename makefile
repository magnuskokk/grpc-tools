.PHONY: all
all: generate

.PHONY: generate
generate:
	generate-grpc
	generate-http
	generate-client
	generate-doc
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
