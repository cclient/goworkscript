#!/bin/sh
cp docker_file/Dockerfile.base ./Dockerfile
docker build -t site_mirror/base --rm ./
rm -f ./Dockerfile