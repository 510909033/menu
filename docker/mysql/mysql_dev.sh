#!/bin/sh

env=dev

docker pull mysql:5.7
docker run -p 5306:3306 --name mysql_$env \
        -v /usr/local/docker/mysql_$env/conf:/etc/mysql \
        -v /usr/local/docker/mysql_$env/logs:/var/log/mysql \
        -v /usr/local/docker/mysql_$env/data:/var/lib/mysql \
        -e MYSQL_ROOT_PASSWORD=LmqI928ZkdPlm\
        -d mysql:5.7
