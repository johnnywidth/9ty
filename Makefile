SERVER_OUT := "env/bin/server"
CLIENT_OUT := "env/bin/client"
API_OUT := "api/*.pb.go"

PKG := "github.com/johnnywidth/9ty"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api server client

all: build-client-image build-server-image

api: api/port_domain.pb.go ## Auto-generate grpc go sources

api/port_domain.pb.go:
	@protoc --go_out=. \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	api/port_domain.proto

build-server: api ## Build the binary file for server
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build-client: api ## Build the binary file for client
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

build-client-image: build-client
	@docker build --no-cache -f env/Dockerfile-client -t 9ty-client:latest .

build-server-image: build-server
	@docker build --no-cache -f env/Dockerfile-server -t 9ty-server:latest .

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
