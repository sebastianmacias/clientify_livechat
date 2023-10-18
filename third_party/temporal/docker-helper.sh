#!/bin/bash

# Path to the compose-manifest.txt file
manifest_file="compose-manifest.txt"

# Validate that the manifest file exists
if [ ! -f "$manifest_file" ]; then
    echo "compose-manifest.txt file does not exist. Exiting."
    exit 1
fi

# Read the docker-compose file name from the manifest file
compose_file=$(cat "$manifest_file")

# Validate that compose_file is not empty and the file exists
if [[ -z "$compose_file" || ! -f "$compose_file" ]]; then
    echo "Invalid or missing docker-compose file specified in compose-manifest.txt: $compose_file"
    exit 1
fi

# Command to --up (start) the docker-compose services
if [ "$1" = "--up" ]; then
    # Use docker-compose up
    docker-compose -f "$compose_file" up
fi

# Command to --logs (display logs) the docker-compose services
if [ "$1" = "--logs" ]; then
    docker-compose -f "$compose_file" logs -f --tail=100
fi

# Command to --halt (stop) the docker-compose services
if [ "$1" = "--halt" ]; then
    docker-compose -f "$compose_file" stop
fi

# Command to --rebuild (force recreate without dependencies and build) the docker-compose services
if [ "$1" = "--rebuild" ]; then
    docker-compose -f "$compose_file" up -d --force-recreate --no-deps --build service
fi

# Command to --destroy (remove containers, networks, volumes, and images)
if [ "$1" = "--destroy" ]; then
    docker-compose -f "$compose_file" down --rmi local -v --remove-orphans
fi

# Validate that one of the supported flags is used
[ -n "$1" -a \( "$1" = "--up" -o "$1" = "--logs"  -o "$1" = "--halt" -o "$1" = "--rebuild" -o "$1" = "--destroy" \) ] \
    || { echo "usage: $0 --up | --logs | --halt | --rebuild | --destroy" >&2; exit 1; }
