.PHONY: generate
generate:
	$(info - Running //go:generate directives) 
	@go generate -x ./...

.PHONY: test
test:
	$(info - Running all go tests)
	@go test -count=1 -v ./...

.PHONY: bench
bench:
	$(info - Running all go benchmarks)
	@go test -run=xxx -bench=. ./...

.PHONY: clean
clean:
	$(info - Removing all generated go files and directories)
	@find . -type f -iname "*.pb.go" -delete
	@find . -type f -iname "*.pb.gw.go" -delete
	@find . -type f -iname "*pb_test.go" -delete
	@find . -type d -name "mocks" -exec rm -rf {} +
	@rm -f statik/statik.go
