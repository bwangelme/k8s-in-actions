#!/usr/bin/env bash

# build
docker build -t bwangel/kubia:latest .

# run
docker run -d -p 8080:8080 --name kubia-container bwangel/kubia:latest

# push
docker push bwangel/kubia:latest
