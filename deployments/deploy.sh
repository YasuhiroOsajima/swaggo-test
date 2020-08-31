# Please use from top of project directory.
# bash deployments/deploy.sh

rm -rf deployments/mysql/data
mkdir deployments/mysql/data

# docker-compose -f deployments/docker-compose.yml up -d
podman-compose -f deployments/docker-compose.yml up -d
