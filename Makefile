
all: main.go
	GOOS=linux GOARCH=mipsle GOMIPS=hardfloat go build -trimpath  -o timelapse main.go
