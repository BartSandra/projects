#!/bin/bash

# Проверка, является ли указанный путь абсолютным путем (первый параметр)
if [[ ! -d "$Absolute_path" ]]; then
    echo "Error! The path must be an absolute path"
    exit 1
fi

# Проверка количества подпапок (второй параметр)
if [[ ! $Number_of_subfolders =~ ^[0-9]+$ ]]; then
    echo "Error! Invalid number of subfolders"
    exit 1
elif [[ $Number_of_subfolders -gt 100 || $Number_of_subfolders -le 0 ]]; then
    echo "Error! Invalid number of folders"
    exit 1
fi

# Проверка количества букв в названии подпапок (третий параметр)
if [[ ${#Folders_name} -gt 7 ]]; then
    echo "Error! The number of letters in folder names should be no more than 7"
    exit 1
elif [[ ! $Folders_name =~ ^[A-Za-z]+$ ]]; then
    echo "Error! Invalid arguments in folder names"
    exit 1
fi

# Проверка количества файлов в подпапках (четвертый параметр)
if [[ ! $Number_of_files =~ ^[0-9]+$ ]]; then
    echo "Error! Invalid argument in number of files"
    exit 1
elif [[ $Number_of_files -gt 100 ]]; then
    echo "Error! Too many files"
    exit 1
fi

# Проверка имени файла и расширения файла (пятый параметр)
export filename=$(echo "$File_name_with_extension" | awk -F. '{print $1}')
export fileExt=$(echo "$File_name_with_extension" | awk -F. '{print $2}')
if [[ ${#filename} -gt 7 || ${#fileExt} -gt 3 || ! $filename =~ ^[A-Za-z]+$ || ! $fileExt =~ ^[A-Za-z]+$ || -z $filename ]]; then
    echo "Error: Invalid argument"
    exit 1
fi

# Проверка размера файла (шестой параметр)
export filesize=$(echo "$File_size" | awk -F"kb" '{print $1}')
if [[ ! $File_size =~ kb$ || ! $filesize =~ ^[0-9]+$ || $filesize -gt 100 || $filesize -le 0 ]]; then
    echo "Error: Invalid size"
    exit 1
fi

bash generator.sh
