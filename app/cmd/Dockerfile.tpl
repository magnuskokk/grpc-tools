# DO NOT EDIT. This file is generated by app/cmd/generate-dockerfiles.sh

# Extend gobuild which has predownloaded dependencies.
FROM grpc-tools/gobuild as build
ENV CGO_ENABLED=0
WORKDIR /build
RUN go build -a -o bin/${name} cmd/${name}/main.go

# Create a single binary image.
FROM scratch as ${name}
COPY --from=build /build/bin/${name} /${name}/${name}
EXPOSE ${port}
CMD ["/${name}/${name}"]