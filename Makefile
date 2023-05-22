include .env

############### GLOBAL VARS ###############
COMPILEDAEMON_PATH=~/go/bin/CompileDaemon # CompileDaemon for hot reload
GO_SERVER=ap-gift-card-server
#############################################
############### LOCAL BUILD #################
#############################################

# dev-mode
.phony: dev
dev: 
	@$(COMPILEDAEMON_PATH) -command="./$(GO_SERVER)"

# local run
.phony: go-run
go-run:
	@go run .
