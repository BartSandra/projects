#!/bin/bash

# запускаем скрипт:
# 1- путь к каталогу (например: ./01), 
# 2- количество подпапок (например: 5), 
# 3- имя папок (например: az), 
# 4- количество файлов (например: 10), 
# 5- имя файла с расширением (например: az.az), 
# 6- размер файла (например: 3kb)

if [[ $# -ne 6 ]]; then
    echo "Error! The script is launched with 6 parameters. Example of running a script:
main.sh /opt/test 4 az 5 az.az 3kb"
    exit 1
fi

# Получаем значения параметров из аргументов командной строки
export Absolute_path=$1
export Number_of_subfolders=$2
export Folders_name=$3
export Number_of_files=$4
export File_name_with_extension=$5
export File_size=$6

bash check.sh