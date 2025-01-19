#!/bin/bash

# Define container name
CONTAINER_NAME="inv-db-standalone"

# Check if the container exists
if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Container '${CONTAINER_NAME}' already exists. Starting it..."
    docker start ${CONTAINER_NAME}
else
    echo "Container '${CONTAINER_NAME}' does not exist. Creating and starting it..."
    docker run -d --name ${CONTAINER_NAME} \
        -e POSTGRES_USER=some-postgres \
        -e POSTGRES_DB=some-postgres \
        -e POSTGRES_PASSWORD=mysecretpassword \
        -p 5432:5432 postgres:16.3
    make mup
fi
