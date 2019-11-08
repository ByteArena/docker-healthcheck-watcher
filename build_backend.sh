#!/bin/sh

docker run --rm -i \
-e GO111MODULE=on \
-e GOOS=linux \
-e GOARCH=amd64 \
-e CGO_ENABLED=0 \
-v $(pwd):/app \
-w /app \
golang:1.12.12-stretch \
go build -mod=vendor -o docker-healthcheck-watcher -ldflags "-X main.version=$DOCKER_TAG -X main.build=$TIME_STAMP -X main.gitCommit=$GIT_COMMIT_SHA" .

