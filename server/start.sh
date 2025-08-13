#!/bin/sh
set -ex

echo "--- DEBUGGING START ---"
echo "--- DEBUGGING START ---"
echo "DATABASE_URL value is: [$DATABASE_URL]"
echo "--- DEBUGGING END ---"

# Retry connecting to the database and running migrations
until /app/goose -dir /app/migrations postgres "$DATABASE_URL" up; do
  echo "goose command failed, retrying in 5 seconds..."
  sleep 5
done

echo "Starting API server"
./server

