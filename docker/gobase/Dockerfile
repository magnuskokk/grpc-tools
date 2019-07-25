FROM golang:1.12.7-alpine
RUN apk add --no-cache git

ENV CGO_ENABLED 0
ENV XDG_CACHE_HOME=/tmp/.cache
ENV APP_PATH /app

ARG USER
ARG UID
ARG GID

RUN addgroup --gid "$GID" "$USER" \
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN mkdir -p ${APP_PATH} \
    && chown -R ${UID}:${GID} ${APP_PATH} \
    && chown -R ${UID}:${GID} /go

WORKDIR ${APP_PATH}

USER ${USER}
