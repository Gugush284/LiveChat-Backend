.PHONY: build_server
buildserver:
	go build -v ./cmd/server


.PHONY: build_run_server_win
build_run_server:
	go build -v ./cmd/server
	./server.exe


.PHONY: build_json
buildjson:
	go build -v ./third_party/json-example/cmd/json-example


.PHONY: build_run_json_win
build_run_json:
	go build -v ./third_party/json-example/cmd/json-example
	./json-example.exe


.PHONY: build_run_json_mac
build_run_json:
	go build -v ./third_party/json-example/cmd/json-example
	./json-example.cmd

.DEFAULT_GOAL := build_server