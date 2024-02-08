#!/bin/bash

# Проверка количества букв в названии подпапок (первый параметр)
if [[ ($Folders_name =~ [0-9]) || ($Folders_name =~ [^A-Za-z]) || (${#Folders_name} -gt 7) ]]; then
    echo "Error! Invalid folder name"
    exit 1
fi

# Проверка имени файла и расширения файла (второй параметр)
export filename="$(echo $Folders_name | awk -F. '{print $1}')"
export fileExt="$(echo $File_name_with_extension | awk -F. '{print $2}')"
if [[ ${#filename} -gt 7 || ${#fileExt} -gt 3 || $filename =~ [^A-Za-z] || $fileExt =~ [^A-Za-z] || ($fileExt == "") || ${#filename} -lt 1 ]]; then
    echo "Error! Invalid files name"
    exit 1
fi

# Проверка размера файла (третий параметр)
export filesize=$(echo "$File_size" | awk -F"Mb" '{print $1}')
if [[ !($File_size =~ Mb$) || ($filesize =~ [^0-9]) || ($filesize -gt 100) || ($filesize -le 0) ]]; then
    echo "Error! Invalid size"
    exit 1
fi

bash file_system_clogging.sh