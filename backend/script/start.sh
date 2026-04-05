#!/bin/bash

APP_NAME="mm-wiki"
APP_DIR=$(cd $(dirname $0); cd ..; pwd)
PID_FILE="$APP_DIR/mm-wiki.pid"
LOG_DIR="$APP_DIR/logs"
CONF_DIR="$APP_DIR/conf"

ENV="prod"
if [ $# -ge 1 ]; then
    ENV=$1
fi

mkdir -p "$LOG_DIR"

if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    if kill -0 "$PID" 2>/dev/null; then
        echo "$APP_NAME is already running (PID: $PID)"
        exit 1
    fi
    rm -f "$PID_FILE"
fi

echo "Starting $APP_NAME with env=$ENV ..."
nohup "$APP_DIR/bin/$APP_NAME" -conf "$CONF_DIR/$ENV/app.yaml" > "$LOG_DIR/stdout.log" 2>&1 &
echo $! > "$PID_FILE"
echo "$APP_NAME started (PID: $!)"
