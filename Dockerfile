FROM scratch
ADD ./docker-healthcheck-watcher /docker-healthcheck-watcher

ENTRYPOINT ["./docker-healthcheck-watcher"]