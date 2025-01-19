#!/bin/bash

# Fail immediately if any command fails
set -e

# Define the Docker image name and tag
IMAGE_NAME="inv-meul-app"
IMAGE_TAG="latest"
DOCKERFILE_PATH="build/Dockerfile"
BUILD_CONTEXT_FROM_DOCKERFILE="."
BUILD_CONTEXT=".."
export DOCKER_BUILDKIT=1

# Paths to secret files
DB_PASSWORD_FILE="db_password.txt"
DB_USER_FILE="db_user.txt"

echo "Building Docker image $IMAGE_NAME:$IMAGE_TAG from Dockerfile $DOCKERFILE_PATH with no cache..."
cd $BUILD_CONTEXT

# Create the secret files with sensitive data
echo "${DB_PASSWORD}" > $DB_PASSWORD_FILE
echo "${DB_USER}" > $DB_USER_FILE

# Build the Docker image with secrets
docker build --secret id=db_password_file,src=$DB_PASSWORD_FILE \
    --secret id=db_user_file,src=$DB_USER_FILE \
    --build-arg BUILD_MODE="release" \
    --build-arg DB_HOST="${DB_HOST}" \
    --build-arg DB_NAME="${DB_NAME}" \
    --build-arg DB_PORT="${DB_PORT}" \
    --build-arg DB_SSL_MODE="${DB_SSL_MODE}" \
    --no-cache --progress=plain -t $IMAGE_NAME:$IMAGE_TAG \
    -f $DOCKERFILE_PATH $BUILD_CONTEXT_FROM_DOCKERFILE

# Clean up the secret files after the build
rm -f $DB_PASSWORD_FILE $DB_USER_FILE

echo "Docker image '${IMAGE_NAME}:${IMAGE_TAG}' built successfully without cache!"
