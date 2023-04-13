.PHONY: buildserver
buildserver:
	go build -v ./cmd/server


.PHONY: runserver
runserver:
	./server.exe


.PHONY: build_run_server
build_run_server:
	go build -v ./cmd/server
	./server.exe


.DEFAULT_GOAL := build_run_server