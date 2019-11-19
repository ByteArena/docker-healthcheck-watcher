#!/bin/sh

echo Building image ${IMAGE_NAME}
docker build --rm -t ${IMAGE_NAME} . 
