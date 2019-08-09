# Used in Travis CI.
.PHONY: vet
vet:
	./tusk go.vet ./...

.PHONY: test
test:
	./tusk go.test ./...

.PHONY: build
build:
	./tusk stack.build
