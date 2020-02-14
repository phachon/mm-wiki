#!/bin/sh

# start go mod
export GO111MODULE=on
# set goproxy
export GOPROXY=https://goproxy.cn

PROJECT_NAME="mm-wiki"
INSTALL_NAME="install"
BUILD_DIR="release"
ROOT_DIR=$(pwd)

# windows .exe
if [ "${GOOS}" = "" ]; then
  UNAME=$( command -v uname)
  case $( "${UNAME}" | tr '[:upper:]' '[:lower:]') in
    msys*|cygwin*|mingw*|nt|win*)
      PROJECT_NAME=${PROJECT_NAME}".exe"
      INSTALL_NAME=${INSTALL_NAME}".exe"
      ;;
  esac
elif [ "${GOOS}" = "windows" ]; then
    PROJECT_NAME=${PROJECT_NAME}".exe"
    INSTALL_NAME=${INSTALL_NAME}".exe"
fi

rm -rf ${BUILD_DIR}

build_app() {
    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/conf
    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/logs
    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/static
    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/views
    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/docs

    /bin/cp -r "${ROOT_DIR}"/conf/default.conf "${ROOT_DIR}"/${BUILD_DIR}/conf/
    /bin/cp -r "${ROOT_DIR}"/conf/template.conf "${ROOT_DIR}"/${BUILD_DIR}/conf/
    /bin/cp -r "${ROOT_DIR}"/scripts/* "${ROOT_DIR}"/${BUILD_DIR}/
    /bin/cp -r "${ROOT_DIR}"/docs/* "${ROOT_DIR}"/${BUILD_DIR}/docs/
    /bin/cp -r "${ROOT_DIR}"/views/* "${ROOT_DIR}"/${BUILD_DIR}/views/
    /bin/cp -r "${ROOT_DIR}"/static/* "${ROOT_DIR}"/${BUILD_DIR}/static/

    /bin/cp -r "${ROOT_DIR}"/CHANGELOG.md "${ROOT_DIR}"/${BUILD_DIR}
    /bin/cp -r "${ROOT_DIR}"/README.md "${ROOT_DIR}"/${BUILD_DIR}
    /bin/cp -r "${ROOT_DIR}"/README_eng.md "${ROOT_DIR}"/${BUILD_DIR}
    /bin/cp -r "${ROOT_DIR}"/LICENSE "${ROOT_DIR}"/${BUILD_DIR}

    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/conf/
    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/logs/
    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/static/
    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/views/
    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/*.sh

    go build -o ${PROJECT_NAME} ./

    if [ -f "${ROOT_DIR}/${PROJECT_NAME}"  ]; then
        mv "${ROOT_DIR}"/${PROJECT_NAME} "${ROOT_DIR}"/${BUILD_DIR}/
    fi
    return
}

build_install() {

    mkdir -p "${ROOT_DIR}"/${BUILD_DIR}/install
    # Todo: 目录切换失败时输出错误信息
    cd "${ROOT_DIR}"/install || exit
    go build -o ${INSTALL_NAME} ./
    chmod -R 755 "${ROOT_DIR}"/${BUILD_DIR}/install/

    if [ -f "${ROOT_DIR}/install/${INSTALL_NAME}"  ]; then
        mv "${ROOT_DIR}"/install/${INSTALL_NAME} "${ROOT_DIR}"/${BUILD_DIR}/install
    fi

    cd ../
    return
}

build() {
    echo ">> MM-Wiki start build!"
    build_app
    build_install
    return
}

build
