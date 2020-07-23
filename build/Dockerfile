FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/catalog-service-claim-operator \
    USER_UID=1001 \
    USER_NAME=catalog-service-claim-operator

# install operator binary
COPY build/_output/bin/catalog-service-claim-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
