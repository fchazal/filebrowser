#!/bin/sh

cd backend
env GOOS=linux GOARCH=arm go build -o ../filebrowser

cd http
~/go/bin/rice append --exec ../../filebrowser