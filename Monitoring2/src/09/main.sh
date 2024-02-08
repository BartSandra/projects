#!/bin/bash

if [[ $# -ne 0 ]]
then
    echo "The script runs without parameters!" >&2
    exit 1
fi

sudo bash start.sh

echo "go to http://localhost:9090/..." >&2

while sleep 3
do
    sudo bash metrics.sh
    sudo cp metrics.html /usr/share/nginx/html/metrics.html
done