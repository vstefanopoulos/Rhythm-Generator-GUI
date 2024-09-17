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

    echo "Deleting source files..."
    rm -f $PACKAGE_PATH/*.go ./main.go ./go.mod ./go.sum
    rm -rf $PACKAGE_PATH

    echo "Source files deleted."
    echo "To run the main program, use $./main.bin instead of \$go run main.go"
    rm -- "$0" &
    
else
    echo "Build failed. Source files remain."
fi

