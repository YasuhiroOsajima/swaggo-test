# Please use from top of project directory.
# bash build/build.sh

# docker build
#docker build -t swaggo-test:0.1 -f build/Dockerfile .
podman build -t swaggo-test:0.1 -f build/Dockerfile
