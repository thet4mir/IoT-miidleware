#!/usr/bin/env bash

echo --- Building my-envoy docker image ---
docker build -t my-envoy:1.1 .


echo --- Running my-envoy docker image ---
docker run --name envoy-proxy -p 8000:8000 -p 9901:9901 --network host my-envoy:1.1