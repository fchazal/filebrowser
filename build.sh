#!/bin/sh

cd backend
env GOOS=linux GOARCH=arm go build -o ../filebrowser

#go get github.com/GeertJohan/go.rice
#go get github.com/GeertJohan/go.rice/rice

cd http
~/go/bin/rice append --exec ../../filebrowser