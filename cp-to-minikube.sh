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
minikube cp .image /image
minikube ssh -- docker rmi -f $IMAGE_NAME 2>/dev/null
minikube ssh -- docker load -i /image
rm .image
