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