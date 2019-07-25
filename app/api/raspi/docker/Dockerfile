FROM golang:1.12.7-alpine AS buildable
RUN apk add --no-cache build-base git
WORKDIR /build
COPY . .
RUN go mod download


FROM buildable AS build
ENV GOOS=linux
RUN go build -a -o bin/raspi-server cmd/raspi-server/main.go
RUN chmod +x bin/raspi-server


FROM alpine:3.10.1
LABEL maintainer=magnuskokk
COPY --from=build /build/bin/raspi-server /raspi-server/raspi-server
EXPOSE 8000
CMD ["/raspi-server/raspi-server"]
