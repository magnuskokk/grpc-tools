.PHONY: generate
generate:
	@go generate -x ./...

.PHONY: test
test:
	@go test -count=1 -v ./...

.PHONY: bench
bench:
	@go test -run=xxx -bench=. ./...

.PHONY: clean
clean:
	$(info - Removing all generated go files and directories)
	@find . -type f -iname "*.pb.go" -delete
	@find . -type f -iname "*.pb.gw.go" -delete
	@find . -type f -iname "*pb_test.go" -delete
	@find . -type d -name "mocks" -exec rm -rf {} +
	@rm -f statik/statik.go

.PHONY: docserver
docserver:
	$(info - Serving documentation at: http://localhost:8000/openapi-ui)
	@go run cmd/docserver/main.go
