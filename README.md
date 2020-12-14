# 9ty

## Before we start
```
 > make api
 > go mod vendor
```

Please, check env/client/ports.json - this file will be send from client to server and saved to datastore, modify it if you want.

## How to run services
```
 > make
 > docker-compose -f ./env/docker-compose.yml up server
 > docker-compose -f ./env/docker-compose.yml up client
 > curl -v -XGET http://localhost:9090/port/AEAUH
```

### What if I want to change ports.json
- stop client docker container
- change env/client/ports.json, no need to re-create docker image as this folder mounted to container
- run `docker-compose -f ./env/docker-compose.yml up client`

### What if I want to erase datastore
- stop server docker container and run `docker-compose -f ./env/docker-compose.yml up server`

## Run tests and linter
```
 > make gen-mock
 > make test
 > make lint
```
