#!/bin/bash
GIT_HASH=$(git rev-parse --short HEAD)
img=robherley/court-herald:$GIT_HASH
echo Building $img:
docker build -t $img .
echo Pushing $img:
docker push $img

while true; do
    read -p "Do you want to update the deployment? (y/n)" yn
    case $yn in
        [Yy]* ) kubectl set image deployment/court-herald court-herald=$img; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y or n.";;
    esac
done
