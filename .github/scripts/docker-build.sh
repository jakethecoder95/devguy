#!/bin/bash
docker build \
  --tag "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:$IMAGE_TAG" \
  --tag "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:$GITHUB_SHA" \
  ../../app

docker push "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:$IMAGE_TAG"
