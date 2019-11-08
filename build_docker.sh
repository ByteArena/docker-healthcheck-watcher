#!/bin/sh
# This file should not need changes

IMAGE_NAME=${REGISTRY}/${REPOSITORY}:${DOCKER_TAG}

echo Building image ${IMAGE_NAME}
docker build -t ${IMAGE_NAME} . || exit 2

echo Logging in to container registry ${REGISTRY} as ${STRATSYS_CR_LOGIN_NAME}
docker login ${REGISTRY} -u ${STRATSYS_CR_LOGIN_NAME} -p ${STRATSYS_CR_LOGIN_KEY}

echo Pushing image ${IMAGE_NAME} to registry
docker push ${IMAGE_NAME}