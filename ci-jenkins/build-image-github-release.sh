#!/usr/bin/env bash

docker build -t cholick/github-release ci/image-github-release
docker push cholick/github-release
