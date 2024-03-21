#!/usr/bin/env bash

#verify if the environment variables are defined
if [[ -z "$IMAGE_NAME" || -z "$IMAGE_VERSION" ]]; then
	echo "please, export the variables IMAGE_NAME and IMAGE_VERSION"
	exit 1
fi

#function to build image
build_image(){
	docker build -t "$IMAGE_NAME:$IMAGE_VERSION" .
}

build_image

