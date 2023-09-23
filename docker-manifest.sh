#!/bin/bash

# Initialize an empty array to store compose files
COMPOSE_FILES=()

# Read the compose-manifest.txt file into the array
# Using a while-loop to read each line from the file into the array
while IFS= read -r line || [ -n "$line" ]; do
    # Append each line read into COMPOSE_FILES array
    COMPOSE_FILES+=(-f "$line")
done < "compose-manifest.txt"

# Convert array to string
# The [*] treats array as single string, separating elements by the first character of IFS variable
COMPOSE_FILES_STR="${COMPOSE_FILES[*]}"

# Debugging statements
echo "COMPOSE_FILES array: ${COMPOSE_FILES[@]}"
echo "COMPOSE_FILES_STR: $COMPOSE_FILES_STR"

if [ "$1" = "--up" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up --no-start"
    eval "docker-compose $COMPOSE_FILES_STR start"
    eval "docker-compose $COMPOSE_FILES_STR exec service bash"
fi

if [ "$1" = "--halt" ]; then
    eval "docker-compose $COMPOSE_FILES_STR stop"
fi

if [ "$1" = "--logs" ]; then
    eval "docker-compose $COMPOSE_FILES_STR logs -f --tail=100"
fi

if [ "$1" = "--rebuild" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up -d --force-recreate --no-deps --build service"
fi

if [ "$1" = "--destroy" ]; then
    eval "docker-compose $COMPOSE_FILES_STR down --rmi local -v --remove-orphans"
fi

[ -n "$1" -a \( "$1" = "--up" -o "$1" = "--logs" -o "$1" = "--halt" -o "$1" = "--rebuild" -o "$1" = "--destroy" \) ] \
    || { echo "usage: $0 --up | --logs | --halt | --rebuild | --destroy" >&2; exit 1; }
