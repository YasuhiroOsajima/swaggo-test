# test app with swaggo

this app is to use with following UI.  
https://github.com/YasuhiroOsajima/react-test.git

This app is depending on docker and docker-compose.  
Please install them in ahead of setup.

## setup

Please set your server's name.

```bash
vim deployments/server.env
```

build docker image.

```bash
bash build/build.sh
```

run app with docker-compose.

```bash
bash deployments/deploy.sh
```
