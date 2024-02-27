set GOARCH=arm64
set GOOS=linux
go build -ldflags="-s -w" -o ./ddns_linux_arm64 ./main.go