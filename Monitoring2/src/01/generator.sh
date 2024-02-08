#!/bin/bash

> logfile.log

foldname=$Folders_name
last_symbol_foldname=${foldname: -1}
fileName=$filename
new_fileName=$fileName
last_symbol_filename=${fileName: -1}

Date=$(date +"%d%m%y")
Date_for_log="DATE = $(date +"%d.%m.%y")"

if [[ ${#foldname} -lt 4 ]]; then
    for (( i=${#foldname}; i<4; i++ )); do
        foldname+="$(echo $last_symbol_foldname)"
    done
fi

if [[ ${#fileName} -lt 4 ]]; then
    for (( i=${#fileName}; i<4; i++ )); do
        fileName+="$last_symbol_filename"
    done
fi

if [[ ${#foldname} -gt 3 ]]; then
    for (( i=1; i<=Number_of_subfolders; i++ )); do
        new_fileName="$fileName"
        if [[ ${#new_fileName} -lt 4 ]]; then
            for (( j=${#new_fileName}; j<4; j++ )); do
                new_fileName+="$last_symbol_filename"
            done
        fi

        mkdir "$Absolute_path/$foldname"_"$Date"
        for (( j=1; j<=Number_of_files; j++ )); do
            free_disk_space="$(df -h / | awk '{print $4}' | tail -n1)"
            if [[ ${free_disk_space: -1} == "M" ]]; then
                exit 1
            fi
            fallocate -l $filesize "$Absolute_path/$foldname"_"$Date/$new_fileName"_"$Date.$fileExt"
            echo "$Absolute_path/$foldname"_"$Date/$new_fileName.$fileExt"_"$Date | $Date_for_log | Size of file = $filesize kb." >> logfile.log
            new_fileName+="$(echo $last_symbol_filename)"
        done
        
        echo "folder - $Absolute_path/$foldname"_"$Date | $Date_for_log" >> logfile.log
        foldname+="$last_symbol_foldname"
    done
fi
