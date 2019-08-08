#!/bin/bash

for service in $(grep -rl 'service' /idl | xargs sed -n 's|.*service \([A-Za-z]*\)API[^.]*|\1|p'); do
    go run generate.go --servicename $service
done
