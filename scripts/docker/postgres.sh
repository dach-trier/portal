#!/bin/bash

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found"
    echo "note: .env file is required to configure the development database"
    exit 1
fi

docker run -d \
    -e POSTGRES_USER=$DB_USER \
    -e POSTGRES_PASSWORD=$DB_PASS \
    -e POSTGRES_DB=$DB_NAME \
    -p $DB_PORT:5432 \
    -v ./volumes/postgres:/var/lib/postgresql/18/docker \
    postgres:18.3-alpine3.22
