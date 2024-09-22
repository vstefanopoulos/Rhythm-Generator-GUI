#!/bin/bash

PACKAGE_PATH="./rhythmgenerator"

if [ ! -f go.mod ]; then
    echo "Error: go.mod not found. Please ensure you're in the correct directory."
    exit 1
fi

echo "Building Go programs..."

go build -o rhythmgenerator.bin $PACKAGE_PATH/*.go

go build -o main.bin ./main.go

if [ $? -eq 0 ]; then
    echo "Build completed successfully!"
    read -p "Do you want to delete source files? (y/n): " response
    if [[ "$response" == "y" || "$responce" == "Y" ]]; then
        echo "Deleting source files..."
        rm -f $PACKAGE_PATH/*.go ./main.go ./go.mod ./go.sum
        rm -rf $PACKAGE_PATH
        echo "Source files deleted."
        rm -- "$0" &
    
    elif [[ "$response" == "n" || "$response" == "N" ]]; then
        echo "Source files were not deleted."
        exit 0
    else
        echo "Invalid response. Exiting without deleting files."
        exit 1
    fi

else
    echo "Build failed. Source files remain."
    exit 1
fi

