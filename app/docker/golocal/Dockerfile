FROM golang:alpine
ENV GO111MODULE=on
ENV CGO_ENABLED=0

ARG USER
ARG UID
ARG GID

RUN apk add build-base git

RUN addgroup --gid "$GID" "$USER" \
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN mkdir -p /app \
    && chown -R ${UID}:${GID} /app \
    && chown -R ${UID}:${GID} /go

WORKDIR /app
USER ${USER}
