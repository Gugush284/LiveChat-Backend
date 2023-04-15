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


.PHONY: buildjson
buildjson:
	go build -v ./third_party/json-example/cmd/json-example


.PHONY: runjson
runjson:
	./json-example.exe


.PHONY: build_run_json
build_run_json:
	go build -v ./third_party/json-example/cmd/json-example
	./json-example.exe


.DEFAULT_GOAL := build_run_server