#!/bin/sh
set -ex

# Retry connecting to the database and running migrations
until /app/goose -dir /app/migrations postgres "$DATABASE_URL" up; do
  echo "goose command failed, retrying in 5 seconds..."
  sleep 5
done

# Start the API server
echo "Starting API server"
./server