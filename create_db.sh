#!/usr/bin/env bash

# Runs PG docker container
# Params:
echo "u"
# * 1 - container name, will be also used as db name, user and password
# * 2 - port (Optional)
function run-pg() {
  if [ -z "$1" ]; then
    echo "specify db name"
    exit 1
  fi
  name=$1
  if [ -z "$2" ]
  then
    port='5432'
  else
    port=$2
  fi

  docker run -d --name=$name -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=$name -e POSTGRES_DB=$name -p 5435:5432 postgres:latest
}

run-pg "ApiDb"