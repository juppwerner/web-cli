@echo off
set GOARCH=amd64

set GOOS=windows
go build

set GOOS=darwin
go build

set GOOS=linux
go build