#!/usr/bin/env sh

set -eu

readonly GOOSE_MIGRATIONS_DIR="/app/migrations"
readonly GOOSE_POSTGRES_CONNECTION_STRING="host=$PG_HOST port=$PG_PORT user=$PG_USER password=$PG_PWD database=$PG_DB_NAME sslmode=disable"

goose -dir="$GOOSE_MIGRATIONS_DIR" postgres "$GOOSE_POSTGRES_CONNECTION_STRING" status -v
goose -dir="$GOOSE_MIGRATIONS_DIR" postgres "$GOOSE_POSTGRES_CONNECTION_STRING" up -v

/app/srv