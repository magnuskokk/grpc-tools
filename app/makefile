.PHONY: generate
generate:
	@go generate -x ./...

.PHONY: clean
clean:
	$(info - Removing all generated go files and directories)
	@rm -rf ./generated
	@find . -type d -name "mocks" -exec rm -rf {} +

.PHONY: test
test:
	@go test -count=1 -v ./...

.PHONY: bench
bench:
	@go test -run=xxx -bench=. ./...

.PHONY: run-echo-server
run-echo-server:
	@go run cmd/echo-server/main.go

# TODO take parameter of which service to serve
.PHONY: run-docserver
run-docserver:
	$(info Running docserver at http://localhost:8000)
	@go run cmd/docserver/main.go
