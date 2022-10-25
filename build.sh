#!/bin/sh
set -e

gf build -a amd64,arm64 -s linux -p temp

# 打包Docker
# docker buildx build --platform linux/amd64 -t hjmcloud/go-oss:latest --push .
# docker buildx build --platform linux/arm64 -t hjmcloud/go-oss:latest --push .
docker build -t hjmcloud/go-oss:latest . 
docker push hjmcloud/go-oss:latest