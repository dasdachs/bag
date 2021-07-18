start: server client

server:
	cd ./server && go run server.go

client:
	cd ./client && yarn start

build:
	docker build --build-arg TARGETOS=linux --build-arg TARGETARCH=arm --build-arg inventory:latest .

build-arm:
	docker build --build-arg TARGETOS=linux --build-arg TARGETARCH=arm --build-arg GOARM=7 inventory:arm-latest .

help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: start client server help build-pi