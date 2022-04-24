#!/bin/bash -e

function parse_conf_value() {
    echo $(grep -e "^$1=" $2 | awk '{split($0,array,"="); print array[2]}' | tr -d '"')
}

function help() {
    echo "$0 /path/to/mm-wiki.conf"
}

if [[ "$1" == "" ]]; then
    help
    exit 1
fi

CONF_FILE=$1
[[ ! -f ${CONF_FILE} ]] && echo "Failed to open ${CONF_FILE}" && exit 1
echo "Config file: ${CONF_FILE}"

BACKUP_NAME=backup-`date '+%Y%m%d-%H%M%S'`
BACKUP_DIR=/tmp/${BACKUP_NAME}
mkdir -p ${BACKUP_DIR}

echo "Backup config..."
cp ${CONF_FILE} ${BACKUP_DIR}/

echo "Backup database..."
DB_DATABASE=$(parse_conf_value name ${CONF_FILE})
DB_USERNAME=$(parse_conf_value user ${CONF_FILE})
DB_PASSWORD=$(parse_conf_value pass ${CONF_FILE})
MYSQL_PWD=${DB_PASSWORD} mysqldump -u${DB_USERNAME} --no-tablespaces ${DB_DATABASE} > ${BACKUP_DIR}/${DB_DATABASE}.sql

echo "Backup documents..."
DOC_DIR=$(parse_conf_value root_dir ${CONF_FILE})
tar czf ${BACKUP_DIR}/documents.tgz -C ${DOC_DIR} .

echo "Pack backup..."
tar czf ${BACKUP_NAME}.tgz -C /tmp/${BACKUP_NAME} .
rm -rf /tmp/${BACKUP_NAME}
echo "Backup completed: ${BACKUP_NAME}.tgz"
