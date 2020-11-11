set GO111MODULE=on
set GOPROXY=https://goproxy.io
go mod init
go mod vendor
go run main.go