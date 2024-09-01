#!/bin/bash
docker buildx build \
  --platform "linux/amd64" \
  --tag "jakethecoder95/devguy:latest" \
  --tag "jakethecoder95/devguy:$GITHUB_SHA" \
  ./app

docker push "jakethecoder95/devguy:latest"
docker push "jakethecoder95/devguy:$GITHUB_SHA"
