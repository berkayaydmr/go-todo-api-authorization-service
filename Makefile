SHELL = /bin/sh

.PHONY: nodemongorun
nodemongorun: $(call print-target)
	nodemon --exec go run main.go --signal SIGTERM;   

.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/credit/ pkg/proto/credit/credit.proto --go_out=plugins=grpc:pkg/proto/credit/

.PHONY: server
server: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server server/main.go
	bin/server

.PHONY: client
client: ## Build and run client.
	go build -race -ldflags "-s -w" -o bin/client client/main.go
	bin/client

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef