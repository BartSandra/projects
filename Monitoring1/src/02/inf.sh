#!/bin/bash

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

echo -e "HOSTNAME = $HOSTNAME"
echo -e "TIMEZONE = $TIMEZONE"
echo -e "USER = $USER"
echo -e "OS = $OS"
echo -e "DATE = $DATE"
echo -e "UPTIME = $UPTIME"
echo -e "UPTIME_SEC = $UPTIME_SEC"
echo -e "IP = $IP"
echo -e "MASK = $MASK"
echo -e "GATEWAY = $GATEWAY"
echo -e "RAM_TOTAL = $RAM_TOTAL"
echo -e "RAM_USED = $RAM_USED"
echo -e "RAM_FREE = $RAM_FREE"
echo -e "SPACE_ROOT = $SPACE_ROOT"
echo -e "SPACE_ROOT_USED = $SPACE_ROOT_USED"
echo -e "SPACE_ROOT_FREE = $SPACE_ROOT_FREE"