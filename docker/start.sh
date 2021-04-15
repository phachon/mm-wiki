#!/usr/bin/env sh

path=/app/mm-wiki
pwd=$PWD
if test $pwd != $path
then
  cd $path
fi

if [ ! -f "conf/mm-wiki.conf" ];then
  echo "文件不存在"
  cp conf/default.conf conf/mm-wiki.conf
  sed -i "s/httpport = 8081/httpport = $HTTP_PORT/g" conf/mm-wiki.conf
  sed -i "s/host=\"127.0.0.1\"/host=\"$DB_HOST\"/g" conf/mm-wiki.conf
  sed -i "s/port=\"3306\"/port=\"$DB_PORT\"/g" conf/mm-wiki.conf
  sed -i "s/name=\"mm_wiki\"/name=\"$DB_NAME\"/g" conf/mm-wiki.conf
  sed -i "s/user=\"root\"/user=\"$DB_USER\"/g" conf/mm-wiki.conf
  sed -i "s/pass=\"123456\"/pass=\"$DB_PASS\"/g" conf/mm-wiki.conf
fi

/app/mm-wiki/mm-wiki --conf /app/mm-wiki/conf/mm-wiki.conf
