export GO111MODULE=on
export GOPROXY=https://goproxy.io
go mod init
go mod vendor
go run main.go