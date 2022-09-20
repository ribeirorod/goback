#!/bin/bash
set -e

# Wait for postgres to be ready
until PGPASSWORD=docker psql -h localhost -U postgres -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER react;
	CREATE DATABASE react;
	GRANT ALL PRIVILEGES ON DATABASE react TO react;
EOSQL


psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    \c react;
    CREATE SCHEMA IF NOT EXISTS fe  
        CREATE TABLE IF NOT EXISTS users (
            id              VARCHAR(255) PRIMARY KEY,
            username        VARCHAR(255) NOT NULL,
            email           VARCHAR(255) NOT NULL,
            password        VARCHAR(255) NOT NULL,
            created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
        AUTHORIZATION react;
EOSQL