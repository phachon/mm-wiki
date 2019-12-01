#!/bin/bash

APP=mm-wiki
INSTALL=install
PID_FILE=../log/${APP}.pid
LOG_FILE=../log/${APP}.log

function install() {
    chmod +x ./${INSTALL}
    ./${INSTALL}
    return 0
}

function check_pid() {
    if [[ -f ${PID_FILE} ]];then
        pid=`cat ${PID_FILE}`
        if [[ -n ${pid} ]]; then
            res=`ps -p ${pid}|grep -v "PID TTY" |wc -l`
            return res
        fi
    fi
    return 0
}

function start() {
    check_pid
    run_res=$?
    if [[ ${run_res} -gt 0 ]];then
        echo -n "${APP} is running already, pid="
        cat ${PID_FILE}
        return 1
    fi
    init_conf $1
    running=$?
    if [[ ${running} -ne 0 ]];then
        echo "${APP} start running init conf failed!"
        return 1
    fi
	chmod +x ./${APP}
    nohup ./${APP}  &> ${PID_FILE} &
    echo $! > ${PID_FILE}
    echo "${PID_FILE} started..., pid=$!"
}

function stop() {
    pid=`cat ${PID_FILE}`
    kill ${pid}
    echo "${APP} stoped..."
}

function restart() {
    pid=`cat ${PID_FILE}`
    kill -USR2 ${pid}
}

function status() {
    check_pid
    run_res=$?
    if [[ ${run_res} -gt 0 ]];then
        echo "start"
    else
        echo "stop"
    fi
}

function help() {
    echo "$0 install|start|stop|restart|status"
}

function pid() {
    cat ${PID_FILE}
}

function init_conf() {
    run_env=${DOCKER_ENV}
    if [[ "$1" != "" ]]; then
        run_env=$1
    fi
    if [[ "$run_env" == "" ]]; then
        echo "run_env is empty!"
        return 1
    fi
    if [[ -d ../conf/${run_env} ]]; then
        echo "run_env is ${run_env}!"
        /bin/cp -r ../conf/${run_env}/* ../conf
        return 0
    fi
    echo "$run_env env conf not found!"
    return 1
}

if [[ "$1" == "" ]]; then
    help
elif [[ "$1" == "install" ]];then
    install
elif [[ "$1" == "stop" ]];then
    stop
elif [[ "$1" == "start" ]];then
    start $2
elif [[ "$1" == "restart" ]];then
    restart
elif [[ "$1" == "status" ]];then
    status
elif [[ "$1" == "pid" ]];then
    pid
else
    help
fi

