#! /bin/bash

cp -R ./templates/ ./bin/

mkdir ./bin/videos

cd bin

nohup ./api &
nohup ./schedule &
nohup ./streamserver &
nohup ./web &

echo "deploy finished"