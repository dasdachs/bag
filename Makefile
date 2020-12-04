.PHONY: all

start: server client

server:
	cd ./backend && go run main.go


client:
	cd ./frontend && yarn start

help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo