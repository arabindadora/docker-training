#!/bin/bash

docker run --name static -p 8080:80 -v $(pwd):/usr/share/nginx/html -d nginx
