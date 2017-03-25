#!/usr/bin/env bash
docker build -t boriska70/bmt-builder -f Dockerfile.build .
docker run -t --name bmt-builder -w /go/src/github.com/boriska70/bmt boriska70/bmt-builder