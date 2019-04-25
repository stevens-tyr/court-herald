#!/bin/bash

echo "Type a tag for this build of court-herald:"
echo -n "> "
read tag

img=robherley/court-herald:$tag
echo Building $img:
docker build -t $tag .
echo Pushing $img:
docker push $tag

while true; do
    read -p "Do you want to update the deployment? (y/n)" yn
    case $yn in
        [Yy]* ) kubectl set image deployment/court-herald court-herald=$img; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer y or n.";;
    esac
done
