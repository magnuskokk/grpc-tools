FROM golang:stretch
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /build
COPY . .
RUN go mod download
