#!/bin/sh

env=online

docker pull mysql:5.7
docker run -p 3307:3306 --name mysql_$env \
        -v /usr/local/docker/mysql_$env/conf:/etc/mysql \
        -v /usr/local/docker/mysql_$env/logs:/var/log/mysql \
        -v /usr/local/docker/mysql_$env/data:/var/lib/mysql \
        -e MYSQL_ROOT_PASSWORD=MlqAzOnTn782N\
        -d mysql:5.7
