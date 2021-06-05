#!/bin/bash

docker build -t cyberchef:ws .
docker run --name cyberchef-ws --net cyberchef -p 80:80 -d cyberchef:ws
