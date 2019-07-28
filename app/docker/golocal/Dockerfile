FROM golang:stretch
ENV GO111MODULE=on
ENV CGO_ENABLED=0

ARG USER
ARG GROUP
ARG UID
ARG GID

# Set up local user.
RUN if grep -q ${GROUP} /etc/group; then \
       groupmod -g ${GID} ${GROUP}; \
    else \
       groupadd -f -g ${GID} ${GROUP}; \
    fi \
    && useradd -m -u ${UID} -g ${GID} ${USER} -s /bin/bash \
    && mkdir -p /app \
    && chown -R ${UID}:${GID} /app \
    && chown -R ${UID}:${GID} /go

WORKDIR /app
USER ${USER}
