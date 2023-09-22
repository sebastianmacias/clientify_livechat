#!/bin/bash

# Read the docker-compose.include file into an array
COMPOSE_FILES=()
while IFS= read -r line; do
    COMPOSE_FILES+=("-f $line")
done < compose-manifest.txt

# Convert array to string
COMPOSE_FILES_STR="${COMPOSE_FILES[*]}"

if [ "$1" = "--up" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up --no-start"
    eval "docker-compose $COMPOSE_FILES_STR start"
    eval "docker-compose $COMPOSE_FILES_STR exec service bash"
fi

if [ "$1" = "--halt" ]; then
    eval "docker-compose $COMPOSE_FILES_STR stop"
fi

if [ "$1" = "--rebuild" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up -d --force-recreate --no-deps --build service"
fi

if [ "$1" = "--destroy" ]; then
    eval "docker-compose $COMPOSE_FILES_STR down --rmi local -v --remove-orphans"
fi

[ -n "$1" -a \( "$1" = "--up" -o "$1" = "--halt" -o "$1" = "--rebuild" -o "$1" = "--destroy" \) ] \
    || { echo "usage: $0 --up | --halt | --rebuild | --destroy" >&2; exit 1; }
