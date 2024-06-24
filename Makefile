# note: call scripts from /scripts

hello:
	echo "Hello"

build:
	go build -o cmd/play/main.go

run:
	go run cmd/play/main.go
