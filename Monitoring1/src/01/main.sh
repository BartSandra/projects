#!/bin/bash

parametr=$1

if [ $# -ne 1 ]; then
    echo "ERROR!!!Enter only 1 parameter!"
    exit 1
fi

if [[ "$1" =~ ^[0-9]+$ ]]; then
        echo "ERROR!!!Enter not numbers!!!!"
        exit 1
    fi

echo "$1"