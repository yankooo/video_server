#! /bin/bash

# Build web and other services
cd /f/GoWorks/src/github.com/yankooo/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd /f/GoWorks/src/github.com/yankooo/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd /f/GoWorks/src/github.com/yankooo/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd /f/GoWorks/src/github.com/yankooo/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

echo "build finished"