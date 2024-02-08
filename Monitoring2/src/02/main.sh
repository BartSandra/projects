#!/bin/bash

# 1- список букв английского алфавита, используемый в названии папок (не более 7 знаков).
# 2- список букв английского алфавита, используемый в имени файла и расширении (не более 7 знаков для имени, не более 3 знаков для расширения). 
# 3- размер файла (в Мегабайтах, но не более 100).

START=$(date +%s%N)
TIMES=$(date +%H:%M)

if [[ $# -ne 3 ]]; then
    echo "Error! The script is run with 3 parameters. An example of running a script:
bash main.sh az az.az 3Mb"
    exit 1
fi
    export Folders_name=$1
    export File_name_with_extension=$2
    export File_size=$3
    
    bash check.sh

    END=$(date +%s%N)
    DIFF=$((($END - $START)/10000000))
    TIMEE=$(date +%H:%M)
    echo "Start time: $TIMES"
    echo "End time: $TIMEE"
    echo "Total running time of the script: $DIFF ms"

    echo "">>logfile.log
    echo "Start time: $TIMES" >>logfile.log
    echo "End time: $TIMEE" >>logfile.log
    echo "Total running time of the script $DIFF ms" >>logfile.log