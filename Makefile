
all: main.go
	GOOS=linux GOARCH=mipsle GOMIPS=hardfloat /usr/local/tinygo/bin/tinygo build -no-debug  -o timelapse main.go
