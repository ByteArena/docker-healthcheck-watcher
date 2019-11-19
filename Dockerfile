FROM golang:1.12.13-alpine3.10 as builder
COPY . /build/
WORKDIR /build
RUN go build -mod=vendor -o docker-healthcheck-watcher

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/docker-healthcheck-watcher /docker-healthcheck-watcher
COPY --from=builder /build/template /template

ENTRYPOINT ["./docker-healthcheck-watcher"]
