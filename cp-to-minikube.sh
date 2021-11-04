#!/bin/zsh

source ~/.zshrc
IMAGE_NAME=$1

if [ -z $IMAGE_NAME ]; then
	echo Usage: cp-to-minikube [IMAGE_NAME]
	exit
fi

docker rmi -f $IMAGE_NAME 2>/dev/null
docker build -t $IMAGE_NAME .
docker save $IMAGE_NAME -o .image
mk cp .image /image
mk ssh -- docker rmi -f $IMAGE_NAME 2>/dev/null
mk ssh -- docker load -i /image
rm .image
