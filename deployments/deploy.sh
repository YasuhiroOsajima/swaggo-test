# Please use from top of project directory.
# bash deployments/deploy.sh

rm -rf deployments/mysql/data
mkdir deployments/mysql/data

sed -i '/^SERVER=/d' deployments/server.env
IPADDRESS=`hostname -i | awk '{print $4}'`
echo "SERVER=${IPADDRESS}" >> deployments/server.env

# docker-compose -f deployments/docker-compose.yml up -d
podman-compose -f deployments/docker-compose.yml up -d
