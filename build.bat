set GOOS=linux
set GOARCH=amd64

go build -ldflags "-w -s" -o ./dist/supervisor-event-listener ./supervisor-event-listener.go
