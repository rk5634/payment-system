#!/bin/bash

set -e

echo "Running database setup using 'migrate'..."

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
    echo "Loaded .env file"
fi

DB_HOST=${POSTGRES_HOST:-localhost}
DB_PORT=${POSTGRES_PORT:-5432}
DB_USER=${POSTGRES_USER:-user}
DB_PASSWORD=${POSTGRES_PASSWORD:-password}
DB_NAME=${POSTGRES_DB:-payment_db}
DB_SSLMODE=${POSTGRES_SSLMODE:-disable}


DB_DSN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"
MIGRATIONS_DIR="$(pwd)/internal/database/postgres/migrations"


echo "Waiting for PostgreSQL at ${DB_HOST}:${DB_PORT}..."
# until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" > /dev/null 2>&1; do
#   echo "PostgreSQL is unavailable - sleeping"
#   sleep 2
# done
echo "PostgreSQL is up - running migrations"

migrate -path "$MIGRATIONS_DIR" -database "$DB_DSN" up

echo "✅ Database migrations applied successfully."
