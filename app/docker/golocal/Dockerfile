FROM golang:alpine
ENV GO111MODULE=on
ARG USER
ARG GROUP
ARG UID
ARG GID

RUN apk add --no-cache build-base git
RUN addgroup --gid "$GID" "$USER" \
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

WORKDIR /app
USER ${USER}
