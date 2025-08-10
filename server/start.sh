#!/bin/sh
set -ex

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until nc -z pg 5432; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 1
done
echo "PostgreSQL is up - executing migrations"

# Run migrations using the compiled goose binary
/app/goose -dir /app/migrations postgres "$DATABASE_URL" up

# Start the API server
echo "Starting API server"
/server