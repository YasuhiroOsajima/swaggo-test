# Please use from top of project directory.
# bash build/build.sh

# go build
go build -o build/server cmd/main.go

# docker build
docker build -t swaggo-test:0.1 -f build/Dockerfile
