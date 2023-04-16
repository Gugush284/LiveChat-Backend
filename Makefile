.PHONY: build_server
buildserver:
	go build -v ./cmd/server


.PHONY: build_run_server
build_run_server:
	go build -v ./cmd/server
	./server


.PHONY: build_json
buildjson:
	go build -v ./third_party/json-example/cmd/json-example


.PHONY: build_run_json
build_run_json:
	go build -v ./third_party/json-example/cmd/json-example
	./json-example

.DEFAULT_GOAL := build_server