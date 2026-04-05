#!/bin/bash

APP_NAME="mm-wiki"
APP_DIR=$(cd $(dirname $0); cd ..; pwd)
PID_FILE="$APP_DIR/mm-wiki.pid"

if [ ! -f "$PID_FILE" ]; then
    echo "$APP_NAME is not running (no PID file found)"
    exit 0
fi

PID=$(cat "$PID_FILE")
if ! kill -0 "$PID" 2>/dev/null; then
    echo "$APP_NAME is not running (stale PID: $PID)"
    rm -f "$PID_FILE"
    exit 0
fi

echo "Stopping $APP_NAME (PID: $PID) ..."
kill "$PID"

for i in $(seq 1 30); do
    if ! kill -0 "$PID" 2>/dev/null; then
        echo "$APP_NAME stopped"
        rm -f "$PID_FILE"
        exit 0
    fi
    sleep 1
done

echo "Force killing $APP_NAME (PID: $PID) ..."
kill -9 "$PID"
rm -f "$PID_FILE"
echo "$APP_NAME force stopped"
