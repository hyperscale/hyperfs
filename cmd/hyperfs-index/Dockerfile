# HyperFS Index Builder
FROM golang:1.11-alpine as index-builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/hyperscale/hyperfs/
RUN echo http://mirror.yandex.ru/mirrors/alpine/v3.6/main > /etc/apk/repositories; \
    echo http://mirror.yandex.ru/mirrors/alpine/v3.6/community >> /etc/apk/repositories
RUN echo "ipv6" >> /etc/modules
RUN apk upgrade --update-cache --available && \
    apk add --update --no-cache ca-certificates curl git make && \
    rm -rf /var/cache/apk/*
COPY ./ .
RUN make target/hyperfs-index

# Application
FROM alpine:3.8
ARG VERSION
ARG VCS_URL
ARG VCS_REF
ARG BUILD_DATE
ENV PORT 8000
ENV DEBUG true
RUN apk add --update --no-cache ca-certificates curl && rm -rf /var/cache/apk/*
WORKDIR /opt/hyperfs/
COPY --from=index-builder /go/src/github.com/hyperscale/hyperfs/target/hyperfs-index .
HEALTHCHECK --interval=5s --timeout=2s CMD curl -f http://localhost:${PORT}/health > /dev/null 2>&1 || exit 1
EXPOSE ${PORT} 7946
VOLUME /var/lib/hyperfs-index
ENTRYPOINT ["/opt/hyperfs/hyperfs-index"]

# Metadata
LABEL org.label-schema.vendor="Hyperscale" \
    org.label-schema.url="https://github.com/hyperscale" \
    org.label-schema.name="HyperFS" \
    org.label-schema.description="Scalable Infinite Storage Server." \
    org.label-schema.version="v${VERSION}" \
    org.label-schema.vcs-url=${VCS_URL} \
    org.label-schema.vcs-ref=${VCS_REF} \
    org.label-schema.build-date=${BUILD_DATE} \
    org.label-schema.docker.schema-version="1.0"
