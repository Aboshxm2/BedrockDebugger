@ECHO OFF

set GOOS=linux
set GOARCH=amd64

go build -o build/BedrockDebugger-linux-amd64

set GOOS=darwin
set GOARCH=amd64

go build -o build/BedrockDebugger-darwin-amd64

set GOOS=windows
set GOARCH=amd64

go build -o build/BedrockDebugger-windows-amd64.exe

set GOOS=windows
set GOARCH=386

go build -o build/BedrockDebugger-windows-386.exe

PAUSE