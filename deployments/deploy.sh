# Please use from top of project directory.
# bash deployments/deploy.sh

docker-compose -f deployments/docker-compose.yml up -d
docker-compose -f deployments/docker-compose.yml exec db \
  mysql -u root -p root -h localhost < ${PWD}/deployments/initialize_db.sh