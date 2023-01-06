#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up -config=sql_migrate.yml

echo "start the app"
exec "$@"
