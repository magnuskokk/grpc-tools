#!/bin/bash

# Parse IDLs and generate GRPC server + gateway
for service in $(grep -rl 'service' /idl | xargs sed -n 's|.*service \([A-Za-z]*\)API[^.]*|\1|p'); do
    go run ./cmd/generate-servers/main.go --servicename $service --gateway
done
