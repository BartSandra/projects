#!/bin/bash

if [[ $1 ]]
then
echo "ERROR!!! Run the script without parameters!"
exit
fi

source inf.sh

echo "Wish to write this in a file? (Y/N) "
read -n 1 answer
if [[ $answer == ["Yy"] ]]; then
            filename=$(date +"%d_%m_%y_%H_%M_%S")
            bash inf.sh > $filename.status
            echo "File saved in the current directory."
        else
            echo "File not saved!"
        fi