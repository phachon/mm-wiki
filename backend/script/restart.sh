#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)

ENV="prod"
if [ $# -ge 1 ]; then
    ENV=$1
fi

"$SCRIPT_DIR/stop.sh"
sleep 2
"$SCRIPT_DIR/start.sh" "$ENV"
