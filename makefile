# Used in Travis CI.
.PHONY: build
build:
	./tusk stack.build

.PHONY: vet
vet:
	./tusk go.vet ./...

.PHONY: test
test:
	./tusk go.test ./...
