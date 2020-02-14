#!/bin/bash
VER="$(grep "SYSTEM_VERSION" global/system.go | tr '"' ' ' | awk '{print $3}')"
RELEASE="release-${VER}"

function show_help_text() {
  echo "Usage: $0 [all|windows|darwin|linux|help|-h|--help] [386|amd64]
  By default, run \"$0\" is same as:
    $0 linux amd64
  run \"$0 {OS_TYPE}\" is same as:
    $0 {OS_TYPE} amd64

Example:
  Build windows 386 Package:
    $0 windows 386
  Build windows amd64 Package:
    $0 windows
  or
    $0 windows amd64
  Build mac Package:
    $0 mac
  Build All OS And All ARCH Package:
    $0 all
  Show Help Text:
    $0 help"
}

function clean() {
  rm -rf "${RELEASE}"
  mkdir "${RELEASE}"
}

function build_pack() {
  echo "Start pack $1 $2..."
  GOOS=$1 GOARCH=$2 ./build.sh
  tar -czf "${RELEASE}/mm-wiki-${VER}-$1-$2.tar.gz" -C release .
}

OS=$(echo "$1" | tr '[:upper:]' '[:lower:]')
ARCH=$(echo "$2" | tr '[:upper:]' '[:lower:]')

case "$OS" in
-h | --help | help)
  show_help_text
  ;;
all)
  clean
  build_pack windows 386
  build_pack windows amd64

  build_pack linux 386
  build_pack linux amd64

  build_pack darwin amd64
  ;;
*)
  clean
  if [ "$OS" != "" ]; then
    if [ "$ARCH" != "" ]; then
      build_pack $OS $ARCH
    else
      build_pack $OS amd64
    fi
  else
    build_pack linux amd64
  fi
  ;;
esac

echo 'END'
