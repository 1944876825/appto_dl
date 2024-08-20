@echo off
set GOOS=linux
set GOARCH=386
go build -o appto_dl -tags netgo -ldflags "-s -w"