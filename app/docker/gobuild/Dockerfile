FROM golang:stretch
ENV GO111MODULE=on

WORKDIR /build
COPY . .
RUN go mod download
