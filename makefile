# Used in Travis CI.
.PHONY: all
all:
	./tusk stack.build

.PHONY: test
test:
	./tusk go.test
