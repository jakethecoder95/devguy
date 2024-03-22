#!/bin/bash
docker buildx buikd \
  --platform "linux/amd64" \
  --tag "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:latest" \
  --tag "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:$GITHUB_SHA" \
  ./app

docker push "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:latest"
docker push "$GAR_ZONE-docker.pkg.dev/$PROJECT_ID/$GAR_REPO/$IMAGE:$GITHUB_SHA"

kubectl apply -f k8s
kubectl set image deployments/devguy-deployment devguy=$IMAGE:$GITHUB_SHA
