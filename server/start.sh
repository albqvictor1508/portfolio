#!/bin/sh
set -ex

# Run migrations using the compiled goose binary
/app/goose -dir /app/migrations postgres "$DATABASE_URL" up

# Start the API server
echo "Starting API server"
./server
