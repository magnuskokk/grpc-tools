FROM grpc-tools/gobuild as build
ENV CGO_ENABLED=0
WORKDIR /build
RUN go build -a -o bin/metrics-server cmd/metrics-server/main.go


FROM scratch as metrics-server
COPY --from=build /build/bin/metrics-server /metrics-server/metrics-server
EXPOSE 8001
CMD ["/metrics-server/metrics-server"]
