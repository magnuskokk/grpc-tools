FROM golang:stretch

ENV GO111MODULE=on

ARG USER
ARG GROUP
ARG UID
ARG GID

RUN if grep -q ${GROUP} /etc/group; then \
       groupmod -g ${GID} ${GROUP}; \
    else \
       groupadd -f -g ${GID} ${GROUP}; \
    fi \
    && useradd -m -u ${UID} -g ${GID} ${USER} -s /bin/bash

RUN mkdir -p /app && \
    chown -R ${UID}:${GID} /app && \
    chown -R ${UID}:${GID} /go

WORKDIR /app

USER ${USER}
