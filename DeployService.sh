#!/bin/bash

WORK_PATH=$(dirname $(readlink -f $0))

PROJECT_NAME=${WORK_PATH##*/}

ENV="local"

docker network ls | grep "shorty_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create shorty_service
    fi

echo "ENV=$ENV">.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

docker-compose up -d