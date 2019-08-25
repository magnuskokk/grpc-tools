FROM grpc-tools/gobuild as build
WORKDIR /build
RUN go build -a -o bin/metrics cmd/metrics/main.go

FROM scratch as metrics
COPY --from=build /build/bin/metrics /metrics/metrics
CMD ["/metrics/metrics"]
