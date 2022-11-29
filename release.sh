#!/bin/bash
# Usage: release-server.sh <version>
set -e

# Check if version is specified
if [ -z "$1" ]; then
    echo "Usage: release.sh <version>"
    exit 1
fi 

# Check if version is valid
if ! [[ "$1" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version: $1, should be in format x.y.z"
    exit 1
fi

# 获取version
version=$1

# 获取 major.minor.patch
major=$(echo $version | cut -d. -f1)
minor=$(echo $version | cut -d. -f2)
patch=$(echo $version | cut -d. -f3)

# Check if version is already released
if git tag | grep -q "v$version"; then
    echo "Version v$version is already released"
    exit 1
fi

# Create tag
git tag -a "v$version" -m "v$version"

# Push tag
git push origin "v$version"




# 检测是否登录docker hub
docker login

# 编译并发布镜像
cd go-oss
docker buildx build --platform linux/amd64,linux/arm64 \
    -t gcslaoli/go-oss:latest \
    -t gcslaoli/go-oss:$major \
    -t gcslaoli/go-oss:$major.$minor \
    -t gcslaoli/go-oss:$major.$minor.$patch \
    --push .