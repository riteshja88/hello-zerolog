all: setup run

run:
	go run hello.go

setup:
	./golang_install
	export GOROOT=/home/ritesh/usr/local/go1.17.6
	export PATH=/home/ritesh/usr/local/go1.17.6/bin:$PATH
	go mod init github.com/riteshja88/hello-zerolog
	go get github.com/rs/zerolog/log
