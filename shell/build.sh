# windows
GOOS=windows GOARCH=386 go build -o bin/win-32-fast-share-clipboard.exe *.go
GOOS=windows GOARCH=amd64 go build -o bin/win-64-fast-share-clipboard.exe *.go
# mac
#GOOS=darwin GOARCH=386 go build -o bin/mac-32-fast-share-clipboard *.go
GOOS=darwin GOARCH=amd64 go build -o bin/mac-64-fast-share-clipboard *.go
# linux
GOOS=linux GOARCH=386 go build -o bin/linux-32-fast-share-clipboard *.go
GOOS=linux GOARCH=amd64 go build -o bin/linux-64-fast-share-clipboard *.go