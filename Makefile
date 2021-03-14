.PHONY: all
./
start: server client

server:
	cd ./server && go run server.go

client:
	cd ./client && yarn start
