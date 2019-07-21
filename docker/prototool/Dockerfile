FROM uber/prototool:1.8.1

ARG user_id
ARG group_id
ARG username
ARG group

RUN if grep -q ${group} /etc/group; then \
       groupmod -g ${group_id} ${group}; \
    else \
       groupadd -f -g ${group_id} ${group}; \
    fi \
    && \
    useradd -m -u ${user_id} -g ${group_id} ${username} -s /bin/bash

USER ${username}

ENTRYPOINT ["/usr/bin/prototool"]
