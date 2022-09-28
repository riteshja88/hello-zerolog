all: setup run

run:
	@echo "Running......"
	@export GOROOT=/home/ritesh/usr/local/go1.17.6
	@/home/ritesh/usr/local/go1.17.6/bin/go run hello.go
	cat logs.txt

setup:
	./golang_install
	export GOROOT=/home/ritesh/usr/local/go1.17.6
	/home/ritesh/usr/local/go1.17.6/bin/go mod init github.com/riteshja88/hello-zerolog
	/home/ritesh/usr/local/go1.17.6/bin/go get github.com/rs/zerolog/log
