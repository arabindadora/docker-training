#!/bin/bash

docker build -t cyberchef:ui .
docker run --name cyberchef-ui --net cyberchef -p 8001:8001 -d cyberchef:ui
