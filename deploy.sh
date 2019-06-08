#! /bin/bash

cp -R ./templates/ ./bin/

mkdir ./bin/videos

cd bin

chmod +x api
chmod +x scheduler
chmod +x streamserver
chmod +x web 

nohup ./api &
nohup ./scheduler &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"
