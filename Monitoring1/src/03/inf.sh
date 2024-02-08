#!/bin/bash

B1=$1; T1=$2; B2=$3; T2=$4;
D="\033[0m"

HOSTNAME=$(hostname)
TIMEZONE=$(cat /etc/timezone | tr "\n" " "; date +"%-:::z")
USER=$USER
OS=$(cat /etc/issue | tr -d "\n\l" | cut -d ' ' -f 1-3)
DATE=$(date +"%-d %B %Y %T")
UPTIME=$(uptime -p | cut -b 4-)
UPTIME_SEC=$(Uptime=$(</proc/uptime); echo ${Uptime%%.*})
IP=$(hostname -I)
MASK=$(ip -4 a | grep global | awk '{print $4}')
GATEWAY=$(ip route | grep default | awk '{print $3}')
RAM_TOTAL=$(echo $(free -m | awk '/Mem/{printf "%.3f GB", $2/1024}'))
RAM_USED=$(echo $(free -m | awk '/Mem/{printf "%.3f GB", $3/1024}'))
RAM_FREE=$(echo $(free -m | awk '/Mem/{printf "%.3f GB", $4/1024}'))
SPACE_ROOT=$(echo $(df / | tail -1 | awk '{printf "%.2f MB", $2/1024}'))
SPACE_ROOT_USED=$(echo $( df / | tail -1 | awk '{printf "%.2f MB", $3/1024}'))
SPACE_ROOT_FREE=$(echo $( df / | tail -1 | awk '{printf "%.2f MB", $4/1024}'))

echo -e "${B1}${T1}HOSTNAME${D} = ${B2}${T2}$(hostname)${D}"
echo -e "${B1}${T1}TIMEZONE${D} = ${B2}${T2}$TIMEZONE${D}"
echo -e "${B1}${T1}USER${D} = ${B2}${T2}$USER${D}"
echo -e "${B1}${T1}OS${D} = ${B2}${T2}$OS${D}"
echo -e "${B1}${T1}DATE${D} = ${B2}${T2}$DATE${D}"
echo -e "${B1}${T1}UPTIME${D} = ${B2}${T2}$UPTIME${D}"
echo -e "${B1}${T1}UPTIME_SEC${D} = ${B2}${T2}$UPTIME_SEC${D}"
echo -e "${B1}${T1}IP${D} = ${B2}${T2}$IP${D}"
echo -e "${B1}${T1}MASK${D} = ${B2}${T2}$MASK${D}"
echo -e "${B1}${T1}GATEWAY${D} = ${B2}${T2}$GATEWAY${D}"
echo -e "${B1}${T1}RAM_TOTAL${D} = ${B2}${T2}$RAM_TOTAL${D}"
echo -e "${B1}${T1}RAM_USED${D} = ${B2}${T2}$RAM_USED${D}"
echo -e "${B1}${T1}RAM_FREE${D} =  ${B2}${T2}$RAM_FREE${D}"
echo -e "${B1}${T1}SPACE_ROOT${D} = ${B2}${T2}$SPACE_ROOT${D}"
echo -e "${B1}${T1}SPACE_ROOT_USED${D} = ${B2}${T2}$SPACE_ROOT_USED${D}"
echo -e "${B1}${T1}SPACE_ROOT_FREE${D} = ${B2}${T2}$SPACE_ROOT_FREE${D}"