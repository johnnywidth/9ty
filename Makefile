SERVER_OUT := "env/bin/server"
CLIENT_OUT := "env/bin/client"

PKG := "github.com/johnnywidth/9ty"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"

## GolangCI-Lint version
GOLANGCI_VERSION=1.26.0
GOLANGCI_COMMIT=6bd10d01fde78697441d9c11e2235f0dbb1e2822

all: build-client-image build-server-image

test:
	go test ./... -cover -coverprofile cover-all.out

gen-mock: 
	go generate ./...

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

lint: bin/golangci-lint ## Run golang-cilint with printing to stdout
	./bin/golangci-lint run --out-format colored-line-number

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCI_COMMIT}/install.sh | BINARY=golangci-lint bash -s -- v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
