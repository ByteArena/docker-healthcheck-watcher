FROM alpine:3.10 as alpine
RUN apk add --no-cache ca-certificates

FROM scratch
COPY ./docker-healthcheck-watcher /docker-healthcheck-watcher
COPY ./template /template
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./docker-healthcheck-watcher"]
