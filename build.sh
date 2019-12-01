#!/bin/sh

# start go mod
export GO111MODULE=on
# set goproxy
export GOPROXY=https://goproxy.cn

PROJECT_NAME="mm-wiki"
INSTALL_NAME="install"
BUILD_DIR="release"
ROOT_DIR=`pwd`

rm -rf ${BUILD_DIR}

build_app() {
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/bin
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/conf
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/logs
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/static
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/views
    mkdir -p ${ROOT_DIR}/${BUILD_DIR}/docs

    /bin/cp -r ${ROOT_DIR}/conf/*.conf ${ROOT_DIR}/${BUILD_DIR}/conf/
    /bin/cp -r ${ROOT_DIR}/scripts/* ${ROOT_DIR}/${BUILD_DIR}/bin/
    /bin/cp -r ${ROOT_DIR}/docs/* ${ROOT_DIR}/${BUILD_DIR}/docs/
    /bin/cp -r ${ROOT_DIR}/views/* ${ROOT_DIR}/${BUILD_DIR}/views/
    /bin/cp -r ${ROOT_DIR}/static/* ${ROOT_DIR}/${BUILD_DIR}/static/

    /bin/cp -r ${ROOT_DIR}/CHANGELOG.md ${ROOT_DIR}/${BUILD_DIR}
    /bin/cp -r ${ROOT_DIR}/README.md ${ROOT_DIR}/${BUILD_DIR}
    /bin/cp -r ${ROOT_DIR}/README_eng.md ${ROOT_DIR}/${BUILD_DIR}
    /bin/cp -r ${ROOT_DIR}/LICENSE ${ROOT_DIR}/${BUILD_DIR}
#    if [[ -d ${ROOT_DIR}/scripts ]]; then
#        /bin/cp -r ${ROOT_DIR}/scripts ${ROOT_DIR}/${BUILD_DIR}/
#    fi

    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/bin/
    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/conf/
    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/logs/
#    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/scripts/
    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/static/
    chmod -R 755 ${ROOT_DIR}/${BUILD_DIR}/views/

    go build -o ${PROJECT_NAME} ./

    if [[ -f "${ROOT_DIR}/${PROJECT_NAME}"  ]]; then
        mv ${ROOT_DIR}/${PROJECT_NAME} ${ROOT_DIR}/${BUILD_DIR}/bin
    fi
    return
}

build_install() {
    cd ${ROOT_DIR}/install

    go build -o ${INSTALL_NAME} ./

    if [[ -f "${ROOT_DIR}/install/${INSTALL_NAME}"  ]]; then
        mv ${ROOT_DIR}/install/${INSTALL_NAME} ${ROOT_DIR}/${BUILD_DIR}/bin
    fi
    return
}

build() {
    echo ">> MM-Wiki start build!"
    build_app
    build_install
    return
}

build
