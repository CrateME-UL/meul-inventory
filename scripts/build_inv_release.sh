#!/bin/sh

# Fail immediately if any command fails
set -e

DB_USER=$1
DB_PASSWORD=$2
BUILD_MODE=$3
DB_HOST=$4
DB_NAME=$5
DB_PORT=$6
DB_SSL_MODE=$7

# Prepare ldflags
LD_FLAGS="-X main.buildMode=${BUILD_MODE}"
LD_FLAGS="${LD_FLAGS} -X main.dbHost=${DB_HOST}"
LD_FLAGS="${LD_FLAGS} -X main.dbUser=${DB_USER}"
LD_FLAGS="${LD_FLAGS} -X main.dbPassword=${DB_PASSWORD}"
LD_FLAGS="${LD_FLAGS} -X main.dbName=${DB_NAME}"
LD_FLAGS="${LD_FLAGS} -X main.dbPort=${DB_PORT}"
LD_FLAGS="${LD_FLAGS} -X main.dbSSLMode=${DB_SSL_MODE}"

# Ensure Go modules are tidy
echo "Running go mod tidy..."
go mod tidy
go install github.com/google/wire/cmd/wire@latest

# Change to the inventory command directory and run wire
echo "Running wire in ${CMD_DIR}/inventory..."
cd ${CMD_DIR}/inventory && wire

# Build the Go application with the specified flags
echo "Building the Go application..."
go build -C ${CMD_DIR}/inventory -ldflags "${LD_FLAGS}" -o ${BUILD_DIR}/inv-meul-app inventory.go wire_gen.go

echo "Build completed successfully!"
