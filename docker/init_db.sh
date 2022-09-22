#!/bin/bash

set -e
set -u 

function create_user_and_database() {
	local database=$1
  local user=$2
	echo "  Creating user and database '$database' and granting access to '$user'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE USER $user CREATEDB PASSWORD '$database';
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $user;
EOSQL
}


# Define Database and Users to create
declare -A access                       # TODO - ENV variables from compose file
access["frontend"]="react"
access["backend"]="go"

for db in "${!access[@]}"; do
  echo "Creating database: $db and user: ${access[$db]}"
  create_user_and_database $db ${access[$db]}
done
