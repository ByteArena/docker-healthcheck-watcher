FROM scratch
ADD ./docker-healthcheck-watcher /docker-healthcheck-watcher
ADD ./template /template

ENTRYPOINT ["./docker-healthcheck-watcher"]