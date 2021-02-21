## Material Icons

https://material.io/resources/icons/?style=outline

## Build FileBrowser

env GOOS=linux GOARCH=arm go build

go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice


cd http
~/go/bin/rice append --exec ../filebrowser

