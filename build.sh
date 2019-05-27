#! /bin/bash

# build web ui

cd ./web
go install
cp /f/GoWorks/bin/web /f/GoWorks/bin/video_server_web_ui/web
cp -R /f/GoWorks/src/github.com/yankooo/video_server/templates /f/GoWorks/bin/video_server_web_ui/
