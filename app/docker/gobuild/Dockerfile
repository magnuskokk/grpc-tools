FROM golang:alpine
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN apk add --no-cache git

WORKDIR /build
COPY . .
RUN go mod download
