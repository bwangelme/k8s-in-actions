#!/usr/bin/env bash

# build
docker build -t bwangel/kubia:simple_v2 .

# run
docker run -d -p 8080:8080 --name kubia-container bwangel/kubia:simple_v2
