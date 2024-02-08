#!/bin/bash

Parametr=$1

if [[ $Parametr =~ [^0-9] ]]; then
    echo "Error: Invalid parametr"
    exit 1
else
    case $Parametr in

    1)
        echo "Input path to logfile"
        read log
        if [ ! -e $log ]; then
        echo "Error: Folder $2 doesn't exsist"
        else
            while IFS= read -r line; do
                if [[ "$line" == *"/"* ]]; then
                    line=$(echo "$line" | /bin/awk '{print $1}')
                    rm -rf $line
                fi
            done < $log
        fi
    ;;

    2)
        echo "Enter the date and time (example: YYYY-MM-DD HH:MM)"
        read -p "Write start date and time: " start
        echo "Enter the date and time (example: YYYY-MM-DD HH:MM)"
        read -p "Write end date and time: " end
        if [[ ! "$start" =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}\ [0-9]{2}:[0-9]{2}$ ]] || [[ ! "$end" =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}\ [0-9]{2}:[0-9]{2}$ ]]; then 
        echo "Error! example: YYYY-MM-DD HH:MM"   
        else
        mapfile -t all_path < <(find / -newermt "$start" ! -newermt "$end" 2>/dev/null)
            for files in "${all_path[@]}"; do 
                if [[ $files =~ _[0-9]{6}$ ]]; then 
                    find $files -exec rm -rf {} +
                fi
            done
        fi
    ;;

    3)
        echo "Enter tha mask: _DDMMYY"
        read mask
        array=(~ /boot /dev /etc /lib /lib32 /lib64 /libx32 /lost+found /media /mnt /opt /proc /root /run /snap /srv /sys /tmp /usr /var)
        for item in ${array[*]}
        do
            cd ~
            cd $item
            rm -rf $(find -type d -name "*${mask}")
        done
    ;;

    *)
    echo "Error!"
    ;;
    esac
fi