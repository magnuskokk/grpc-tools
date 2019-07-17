.PHONY: generate
generate:
	@go generate -x ./...

.PHONY: clean
clean:
	$(info - Removing all generated go files and directories)
	@find . -type f -iname "*.pb.go" -delete
	@find . -type f -iname "*.pb.gw.go" -delete
	@find . -type f -iname "*pb_test.go" -delete
	@find . -type f -iname "statik.go" -delete
	@find . -type d -name "mocks" -exec rm -rf {} +

.PHONY: test
test:
	@go test -count=1 -v ./...

.PHONY: bench
bench:
	@go test -run=xxx -bench=. ./...

.PHONY: run-heartbeat-server
run-heartbeat-server:
	@go run cmd/heartbeat-server/main.go

# TODO take parameter of which service to serve
.PHONY: run-docserver
run-docserver:
	$(info Running docserver at http://localhost:8000)
	@go run cmd/docserver/main.go

