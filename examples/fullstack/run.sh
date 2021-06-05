#!/bin/bash

# create a network for this app
docker network create cyberchef

# create and run the services
cd ./api && ./run.sh
cd ../ui && ./run.sh
cd ../ws && ./run.sh
