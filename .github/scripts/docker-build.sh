#!/bin/bash
docker buildx build \
  --platform "linux/amd64" \
  --tag "$DOCKERHUB_USERNAME/$IMAGE:latest" \
  --tag "$DOCKERHUB_USERNAME/$IMAGE:$GITHUB_SHA" \
  ./app

docker push "$DOCKERHUB_USERNAME/$IMAGE:latest"
docker push "$DOCKERHUB_USERNAME/$IMAGE:$GITHUB_SHA"

kubectl apply -f k8s
kubectl set image deployments/devguy-deployment devguy=$IMAGE:$GITHUB_SHA
