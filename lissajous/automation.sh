#!/usr/bin/env bash

#verify if the environment variables are defined
if [[ -z "$IMAGE_NAME" || -z "$IMAGE_VERSION" ]]; then
	echo "please, export the variables IMAGE_NAME and IMAGE_VERSION"
	exit 1
fi

#function to build image
build_image(){
	echo -e "building image, please wait..."
	docker build -t "$IMAGE_NAME:$IMAGE_VERSION" .
	sleep 2
	echo -e "image successfully built."
}

#function to push image
push_image() {
    if [[ -z "$DOCKER_USERNAME" ]]; then
        echo "please, export the variable DOCKER_USERNAME"
        exit 1
    fi
    docker tag "$IMAGE_NAME:$IMAGE_VERSION" "$DOCKER_USERNAME/$IMAGE_NAME:$IMAGE_VERSION"
    docker push "$DOCKER_USERNAME/$IMAGE_NAME:$IMAGE_VERSION"
}


#function to run image
run_image() {
    if [[ -z "$HOST_PORT" || -z "$CONTAINER_PORT" ]]; then
        echo "please, export the variables HOST_PORT and CONTAINER_PORT to run the image."
        exit 1
    fi
    docker run -p "$HOST_PORT:$CONTAINER_PORT" "$IMAGE_NAME:$IMAGE_VERSION"
}

#function to show options
show_options() {
echo -e "
Options:
1 - Build image
2 - Push image
3 - Run image
4 - Exit
"
}

#loop to proccess the informations
while true; do
    show_options
    read -p "choose an option: " opt
    case $opt in
        1)
            build_image
            ;;
        2)
            push_image
            ;;
        3)
            run_image
            ;;
        4)
            echo "exit."
            break
            ;;
        *)
            echo "invalid option"
            ;;
    esac
done