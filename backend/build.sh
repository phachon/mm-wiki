#!/bin/sh

export GO111MODULE=on

APP_NAME="mm-wiki"
TARGET="release"
ROOT_DIR=`pwd`

rm -rf $TARGET

ENV="prod"
if [ $# -ge 1 ] ;then
    ENV=$1
fi

# windows .exe
if [ "${GOOS}" = "" ]; then
  UNAME=$( command -v uname)
  case $( "${UNAME}" | tr '[:upper:]' '[:lower:]') in
    msys*|cygwin*|mingw*|nt|win*)
      APP_NAME=${APP_NAME}".exe"
      ;;
  esac
elif [ "${GOOS}" = "windows" ]; then
    APP_NAME=${APP_NAME}".exe"
fi

build() {
    mkdir -p "$ROOT_DIR"/$TARGET/bin
    mkdir -p "$ROOT_DIR"/$TARGET/logs
    mkdir -p "$ROOT_DIR"/$TARGET/script
    mkdir -p "$ROOT_DIR"/$TARGET/conf

    chmod 755 "$ROOT_DIR"/$TARGET/bin/

    go build -o ${APP_NAME} ./

    if [ -f "$ROOT_DIR/$APP_NAME"  ]; then
        mv "$ROOT_DIR"/$APP_NAME "$ROOT_DIR"/$TARGET/bin
    fi
    /bin/cp -r "$ROOT_DIR"/conf/"$ENV"/* "$ROOT_DIR"/$TARGET/conf/
    /bin/cp -r "$ROOT_DIR"/script/* "$ROOT_DIR"/$TARGET/script/
    return
}

build