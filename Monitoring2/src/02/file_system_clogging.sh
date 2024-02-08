#!/bin/bash

> logfile.log

folder_name="$(compgen -d / | shuf -n1)"
foldname=$Folders_name
last_symbol_foldname=${foldname: -1}
fileName=$filename
new_fileName=$fileName
last_symbol_filename=${fileName: -1}

Date="$(date +"%d%m%y")"
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

countOfFolders=100
for (( i=1; i<=$countOfFolders; i++ )); do
    folder_name="$(compgen -d / | shuf -n1)"
    filesCounter="$(shuf -i 50-100 -n1)"
    if [[ $folder_name == "/bin" || $folder_name == "/sbin" || $folder_name == "/proc" || $folder_name == "/sys" ]]; then
        countOfFolders+="$(echo $countOfFolders+1)"
        continue
    fi

    new_fileName="$fileName"
        if [[ ${#new_fileName} -lt 4 ]]; then
            for (( j=${#new_fileName}; j<4; j++ )); do
                new_fileName+="$last_symbol_filename"
            done
        fi

    mkdir "$folder_name/"$foldname"_"$Date"" 2>/dev/null
    for (( j=1; j<=${filesCounter}; j++)); do
        free_disk_space="$(df -h / | awk '{print $4}' | tail -n1)"
        if [[ ${free_disk_space: -1} == "M" ]]; then
            exit 1
        fi
        fallocate -l $filesize"M" ""$folder_name"/"$foldname"_"$Date"/"$new_fileName"_"$Date"."$fileExt"" 2>/dev/null
        echo ""$folder_name"/"$foldname"_"$Date"/"$new_fileName"_"$Date"."$fileExt" | "$Date_for_log" | Size of file = $filesize Mb.">>logfile.log
        new_fileName+="$(echo $last_symbol_filename)"
    done
    echo "folder - "$folder_name"/"$foldname"_"$Date" | "$Date_for_log"">>logfile.log
    foldname+="$(echo $last_symbol_foldname)"
done
