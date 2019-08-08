.PHONY: all
all:
	sh install-tusk.sh
	tusk stack.build

.PHONY: test
test:
	tusk go.test
