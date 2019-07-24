FROM golang:1.12.7-alpine AS buildable
RUN apk add --no-cache build-base git
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


FROM buildable AS build
ENV GOOS=linux

RUN go build -a -o bin/echo-server cmd/echo-server/main.go
RUN chmod +x bin/echo-server


FROM alpine:3.10.1
LABEL maintainer=magnuskokk
COPY --from=build /build/bin/echo-server /echo-server/echo-server
EXPOSE 8000
CMD ["/echo-server/echo-server"]
