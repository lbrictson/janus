#!/bin/bash
tag=`git describe --exact-match --tags`
if [ -z "$tag" ]; then
  echo "No tag found"
  exit 1
fi
docker build -t janus:$tag .
docker build -t janus:$tag-slim -f slim.Dockerfile .
docker tag janus:$tag lbrictson/janus:$tag
docker tag janus:$tag-slim lbrictson/janus:$tag-slim
docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
docker push lbrictson/janus:$tag
docker push lbrictson/janus:$tag-slim