.PHONY: all
all: generate

.PHONY: generate
generate:
	docker-compose -f docker-compose.tools.yml run prototool prototool generate
	docker-compose -f docker-compose.tools.yml run prototool chown -R $(shell id -u):$(shell id -g) /work/app/generated

.PHONY: lint
lint:
	docker-compose -f docker-compose.tools.yml run prototool prototool lint 

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
	rm -rf .gen

.PHONY: sudoclean
sudoclean: clean
	$(info - Force clean with .direnv removal)
	@sudo rm -rf ./.direnv
	docker system prune --volumes -a
