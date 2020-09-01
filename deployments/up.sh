SERVER_IP=`hostname -i | awk '{print $4}'`
podman run \
	--name swaggo-test \
	--publish 3389:3389 \
	-e SERVER=${SERVER_IP} \
	-e PORT=3389 \
	-e DBSERVER=127.0.0.1 \
	-e DBPORT=3306 \
	localhost/swaggo-test:0.1
