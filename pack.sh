#!/bin/bash
VER="$(grep "SYSTEM_VERSION" global/system.go | tr '"' ' ' | awk '{print $3}')"
RELEASE="release-${VER}"

rm -rf "${RELEASE}"
mkdir "${RELEASE}"

function build_pack() {
  echo "Start pack $1 $2..."
  GOOS=$1 GOARCH=$2 ./build.sh
  tar -czf "${RELEASE}/mm-wiki-${VER}-$1-$2.tar.gz" -C release .
}

OS=$(echo $1 | tr '[:upper:]' '[:lower:]')
ARCH=$(echo $2 | tr '[:upper:]' '[:lower:]')

if [ "$OS" = "all" ]; then
  build_pack windows 386
  build_pack windows amd64

  build_pack linux 386
  build_pack linux amd64

  build_pack darwin amd64
elif [ "$OS" != "" ]; then
  if [ "$ARCH" != "" ]; then
    build_pack $OS $ARCH
  else
    build_pack $OS amd64
  fi
else
  build_pack linux amd64
fi

echo 'END'
