WORKSPACE ?= $(shell pwd)
include ${WORKSPACE}/.env
export $(shell sed 's/=.*//' .env)
export WORKSPACE
server-build:
	go build -o ${WORKSPACE}/bin/server ${WORKSPACE}/cmd/server/main.go
server-run:
	${WORKSPACE}/bin/server
build-and-run:
	go build -o ${WORKSPACE}/bin/server ${WORKSPACE}/cmd/server/main.go
	${WORKSPACE}/bin/server