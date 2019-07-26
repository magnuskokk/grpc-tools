FROM golang:alpine
ENV GO111MODULE=on
RUN apk add --no-cache build-base git

WORKDIR /build
COPY . .
RUN go mod download
