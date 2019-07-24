.PHONY: all
all: generate

.PHONY: clean
clean:
	$(MAKE) -C app clean
	$(MAKE) -C frontend clean
	$(MAKE) -C swagger clean

.PHONY: envclean
envclean: clean
	@sudo rm -rf ./.direnv
	docker system prune --volumes -a

user != id -u

.PHONY: generate
generate:
	docker-compose -f docker-compose.tools.yml run -u $(user) prototool \
		prototool generate

	$(MAKE) -C app generate

.PHONY: listlinters
listlinters:
	docker-compose -f docker-compose.tools.yml run -u $(user) prototool \
		prototool lint --list-lint-group uber2

.PHONY: servedoc
servedoc:
	docker-compose -f docker-compose.tools.yml run -p 8080:8080 swagger


.PHONY: test
test:
	$(MAKE) -C app test

