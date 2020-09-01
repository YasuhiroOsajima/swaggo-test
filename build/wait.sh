#!/bin/sh

BIN_PATH=/swaggo-test/build/server
MYSQL_SERVER=mysql_host

until mysqladmin ping -h ${MYSQL_SERVER} --silent; do
  echo 'waiting for mysql server becomes connectable.'
  sleep 2
done

exec ${BIN_PATH}