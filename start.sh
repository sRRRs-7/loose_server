#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up -config=/app/sql_migrate.yml -env="production"

echo "start the app"
exec "$@"
