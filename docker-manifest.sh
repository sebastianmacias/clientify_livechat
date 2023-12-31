#!/bin/bash

# Path to the compose-manifest.txt file
manifest_file="compose-manifest.txt"

# Validate that the manifest file exists
if [ ! -f "$manifest_file" ]; then
    echo "compose-manifest.txt file does not exist. Exiting."
    exit 1
fi

# Initialize an empty array to store compose files
COMPOSE_FILES=()

# Read the compose-manifest.txt file into the array
# Using a while-loop to read each line from the file into the array
while IFS= read -r line || [ -n "$line" ]; do
    # Validate that the Docker Compose file specified in the line exists
    if [[ -z "$line" || ! -f "$line" ]]; then
        echo "Invalid or missing docker-compose file specified in compose-manifest.txt: $line"
        exit 1
    fi
    # Append each line read into COMPOSE_FILES array
    COMPOSE_FILES+=(-f "$line")
done < "$manifest_file"

# Convert array to string
# The [*] treats the array as a single string, separating elements by the first character of IFS variable
COMPOSE_FILES_STR="${COMPOSE_FILES[*]}"

# Debugging statements
echo "COMPOSE_FILES array: ${COMPOSE_FILES[@]}"
echo "COMPOSE_FILES_STR: $COMPOSE_FILES_STR"

# Command to --up (start) the docker-compose services
if [ "$1" = "--up" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up --no-start"
    eval "docker-compose $COMPOSE_FILES_STR start"
    eval "docker-compose $COMPOSE_FILES_STR exec service bash"
fi

# Command to --logs (display logs) the docker-compose services
if [ "$1" = "--logs" ]; then
    eval "docker-compose $COMPOSE_FILES_STR logs -f --tail=100"
fi

# Command to --halt (stop) the docker-compose services
if [ "$1" = "--halt" ]; then
    eval "docker-compose $COMPOSE_FILES_STR stop"
fi

# Command to --rebuild (force recreate without dependencies and build) the docker-compose services
if [ "$1" = "--rebuild" ]; then
    eval "docker-compose $COMPOSE_FILES_STR up -d --force-recreate --no-deps --build service"
fi

# Command to --destroy (remove containers, networks, volumes, and images)
if [ "$1" = "--destroy" ]; then
    eval "docker-compose $COMPOSE_FILES_STR down --rmi local -v --remove-orphans"
fi

# Validate the command-line argument
[ -n "$1" -a \( "$1" = "--up" -o "$1" = "--logs" -o "$1" = "--halt" -o "$1" = "--rebuild" -o "$1" = "--destroy" \) ] \
    || { echo "usage: $0 --up | --logs | --halt | --rebuild | --destroy" >&2; exit 1; }
