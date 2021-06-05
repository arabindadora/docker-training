#!/bin/bash
# utility script to build and publish a docker image to a registry

# configure registry and image path
REGISTRY=registry.gitlab.com
NAMESPACE=akdora
REPO=docker-training
IMAGE_NAME=cyberchef-api
IMAGE_TAG=latest

# login to the registry
docker login $REGISTRY

# build the docker image
docker build -t $REGISTRY/$NAMESPACE/$REPO/$IMAGE_NAME:$IMAGE_TAG ./api

# publish/upload the image
docker push $REGISTRY/$NAMESPACE/$REPO/$IMAGE_NAME:$IMAGE_TAG
