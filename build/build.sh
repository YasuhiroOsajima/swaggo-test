# Please use from top of project directory.
# bash build/build.sh

# go build
docker build -t build-swaggo-test:0.1 -f build/Dockerfile.build

# docker build
docker build -t swaggo-test:0.1 -f build/Dockerfile
