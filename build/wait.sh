#!/bin/sh

BIN_PATH=/swaggo-test/build/server
MYSQL_SERVER=mysql_host

for i in `seq 0 20`
do
  mysqladmin ping -h ${MYSQL_SERVER} --silent
  if [ $? -eq 0 ]; then
    echo 'mysql server became connectable.'
    break
  fi

  echo 'waiting for mysql server becomes connectable.'
  sleep 2
done

exec ${BIN_PATH}