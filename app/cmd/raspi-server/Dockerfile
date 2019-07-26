# DO NOT EDIT. This file is generated.
FROM grpc-tools/gobuild as build
ENV CGO_ENABLED=0
WORKDIR /build
RUN go build -a -o bin/raspi-server cmd/raspi-server/main.go


FROM scratch as raspi-server
COPY --from=build /build/bin/raspi-server /raspi-server/raspi-server
EXPOSE 8001
CMD ["/raspi-server/raspi-server"]
