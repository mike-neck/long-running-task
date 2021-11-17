.PHONY: build
build:
	go build -o build/long-running-task ./main.go

.PHONY:run
run:build
	build/long-running-task -each 10 -until 300
