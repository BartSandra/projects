#!/bin/bash

if [[ $# != 1 ]]; then
    echo "Error! The script is run with 1 parameter"
else
    bash clean.sh $1
fi