#* GET FILE ENV
include .env
export $(shell sed 's/=.*//' .env)


################# TODO: GOLANG #################
start:
	go run ./cmd/server/main.go

