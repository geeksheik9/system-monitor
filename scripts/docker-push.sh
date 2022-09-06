#!/usr/bin/env bash

. ./scripts/version.sh

name="system-monitor"
version="$system_monitor_version"
repo="repository"
namespace="docker"
username="geeksheik9"
versionedImageName="$name:$version"
taggedImage="$username/$versionedImageName"

echo "version $version"

#--platform linux/arm64/v8 is to build the image for ubuntu server 22.04
echo "Building ${taggedImage}"
docker build --platform linux/arm64/v8 --no-cache -t ${taggedImage} --build-arg VERSION=${version} .

echo "Please login to docker hub"
docker logout
docker login --username=$username

echo "pushing image ${taggedImage}"
docker push ${taggedImage}