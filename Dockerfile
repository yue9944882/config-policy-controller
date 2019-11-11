FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/multicloud-operators-policy-controller \
    USER_UID=1001 \
    USER_NAME=multicloud-operators-policy-controller

# install operator binary
COPY build/_output/bin/multicloud-operators-policy-controller ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
