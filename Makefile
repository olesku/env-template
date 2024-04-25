all: main.go
	GOOS=linux go build -ldflags "-w -s"
