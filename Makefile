.PHONY: cmd

export VERSION := $(shell cat VERSION)
export BUILD_TIME := $(shell date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
export GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || true)

export BIN_PATH := $(shell pwd)/bin

cmd:
	go build \
		-ldflags "-X \"github.com/nwpc-oper/nmc-typhoon-db-client/cli/cmd.Version=${VERSION}\" \
        -X \"github.com/nwpc-oper/nmc-typhoon-db-client/cli/cmd.BuildTime=${BUILD_TIME}\" \
        -X \"github.com/nwpc-oper/nmc-typhoon-db-client/cli/cmd.GitCommit=${GIT_COMMIT}\" " \
		-o ${BIN_PATH}/nmc_typhoon_db_client \
		cli/main.go