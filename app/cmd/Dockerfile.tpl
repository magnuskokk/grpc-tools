# DO NOT EDIT. This file is generated.
FROM grpc-tools/gobuild as build
ENV CGO_ENABLED=0
WORKDIR /build
RUN go build -a -o bin/${name} cmd/${name}/main.go


FROM scratch as ${name}
COPY --from=build /build/bin/${name} /${name}/${name}
EXPOSE ${port}
CMD ["/${name}/${name}"]
