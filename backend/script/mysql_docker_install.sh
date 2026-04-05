#!/bin/bash

# Install MySQL 8.0 using Docker for local development
# Usage: ./mysql_docker_install.sh

CONTAINER_NAME="mm-wiki-mysql"
MYSQL_PORT=3306
MYSQL_ROOT_PASSWORD="123456"
MYSQL_DATABASE="mm_wiki2"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Error: Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if container already exists
if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Container '$CONTAINER_NAME' already exists."
    STATUS=$(docker inspect -f '{{.State.Running}}' "$CONTAINER_NAME")
    if [ "$STATUS" = "true" ]; then
        echo "Container is already running."
    else
        echo "Starting existing container..."
        docker start "$CONTAINER_NAME"
    fi
    exit 0
fi

echo "Creating MySQL container: $CONTAINER_NAME ..."
docker run -d \
    --name "$CONTAINER_NAME" \
    -p ${MYSQL_PORT}:3306 \
    -e MYSQL_ROOT_PASSWORD="$MYSQL_ROOT_PASSWORD" \
    -e MYSQL_DATABASE="$MYSQL_DATABASE" \
    -e MYSQL_CHARSET=utf8mb4 \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_general_ci \
    mysql:8.0

echo "Waiting for MySQL to initialize..."
sleep 15

echo "MySQL container '$CONTAINER_NAME' is ready."
echo "  Host: 127.0.0.1"
echo "  Port: $MYSQL_PORT"
echo "  User: root"
echo "  Password: $MYSQL_ROOT_PASSWORD"
echo "  Database: $MYSQL_DATABASE"
echo ""
echo "To import the schema, run:"
echo "  mysql -h 127.0.0.1 -u root -p${MYSQL_ROOT_PASSWORD} ${MYSQL_DATABASE} < docs/database/mm_wiki2.sql"
