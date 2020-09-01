# !/bin/.sh

BIN_PATH=$1
MYSQL_SERVER=127.0.0.1

until mysqladmin ping -h ${MYSQL_SERVER} --silent; do
  echo 'waiting for mysql server becomes connectable.'
  sleep 2
done

exec ${BIN_PATH}