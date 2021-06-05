#!/bin/bash

docker build -t cyberchef:api .
docker run --name cyberchef-api --net cyberchef -p 8000:8000 -d cyberchef:api
