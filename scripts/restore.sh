#!/bin/bash -e

function parse_conf_value() {
    echo $(grep -e "^$1=" $2 | awk '{split($0,array,"="); print array[2]}' | tr -d '"')
}

function help() {
    echo "$0 backup-date-time.tgz"
}

if [[ "$1" == "" ]]; then
    help
    exit 1
fi

ROOT_DIR=$(cd "$(dirname "$0")";pwd)
BACKUP_FILEPATH=$1
BACKUP_FILENAME=$(basename ${BACKUP_FILEPATH})
BACKUP_NAME=${BACKUP_FILENAME%.*}
BACKUP_DIR=/tmp/${BACKUP_NAME}
mkdir -p ${BACKUP_DIR}

echo "Restore backup: ${BACKUP_FILEPATH}"
tar xf ${BACKUP_FILEPATH} -C ${BACKUP_DIR}

echo "Restore config..."
CONF_FILE=${BACKUP_DIR}/mm-wiki.conf
if [[ -f ${ROOT_DIR}/conf/mm-wiki.conf ]]; then
    mv ${ROOT_DIR}/conf/mm-wiki.conf ${ROOT_DIR}/conf/mm-wiki.conf.save
fi
cp ${BACKUP_DIR}/mm-wiki.conf ${ROOT_DIR}/conf/mm-wiki.conf

echo "Restore database..."
DB_DATABASE=$(parse_conf_value name ${CONF_FILE})
DB_USERNAME=$(parse_conf_value user ${CONF_FILE})
DB_PASSWORD=$(parse_conf_value pass ${CONF_FILE})
MYSQL_PWD=${DB_PASSWORD} mysql -u${DB_USERNAME} ${DB_DATABASE} < ${BACKUP_DIR}/${DB_DATABASE}.sql

echo "Restore documents..."
DOC_DIR=$(parse_conf_value root_dir ${CONF_FILE})
if [[ -d ${DOC_DIR} ]]; then
	[[ -d ${DOC_DIR}.save ]] && rm -rf ${DOC_DIR}.save
    mv ${DOC_DIR} ${DOC_DIR}.save
fi
mkdir -p ${DOC_DIR}
tar xf ${BACKUP_DIR}/documents.tgz -C ${DOC_DIR}

rm -rf ${BACKUP_DIR}
echo "Restore completed."
