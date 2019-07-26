FROM grpc-tools/gobuild as build
ENV CGO_ENABLED=0
WORKDIR /build
RUN go build -a -o bin/echo-server cmd/echo-server/main.go


FROM scratch as echo-server
COPY --from=build /build/bin/echo-server /echo-server/echo-server
EXPOSE 8000
CMD ["/echo-server/echo-server"]
