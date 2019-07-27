#!/bin/bash

# Render Dockerfiles for servers
for d in $DEVROOT/app/cmd/*-server/ ; do
    name=$(basename $d)
    port=$(sed -n 's/var\sserveAddr\s=\s\".*:\([0-9]*\)\"/\1/p' $DEVROOT/app/cmd/$name/main.go)
    sed \
        -e "s/\${name}/$name/gi" \
        -e "s/\${port}/$port/gi" \
        $DEVROOT/app/cmd/Dockerfile.tpl \
        > $DEVROOT/app/cmd/$name/Dockerfile
done

