#!/bin/bash

# Background
BBLACK="\033[40m"; BRED="\033[41m"; BGREEN="\033[42m"
BBLUE="\033[44m"; BPURPLE="\033[45m"; BWHITE="\033[47m"

# Foreground
TBLACK="\033[30m"; TRED="\033[31m"; TGREEN="\033[32m"
TBLUE="\033[34m"; TPURPLE="\033[35m"; TWHITE="\033[37m"

case "$2" in
    "0") printf "$TBLACK" ;;
    "1") printf "$TWHITE" ;;
    "2") printf "$TRED" ;;
    "3") printf "$TGREEN" ;;
    "4") printf "$TBLUE" ;;
    "5") printf "$TPURPLE" ;;
    "6") printf "$TBLACK" ;;
esac

case "$1" in
    "0") printf "$BBLACK" ;;
    "1") printf "$BWHITE" ;;
    "2") printf "$BRED" ;;
    "3") printf "$BGREEN" ;;
    "4") printf "$BBLUE" ;;
    "5") printf "$BPURPLE" ;;
    "6") printf "$BBLACK" ;;
esac

printf " "

case "$4" in
    "0") printf "$TBLACK" ;;
    "1") printf "$TWHITE" ;;
    "2") printf "$TRED" ;;
    "3") printf "$TGREEN" ;;
    "4") printf "$TBLUE" ;;
    "5") printf "$TPURPLE" ;;
    "6") printf "$TBLACK" ;;
esac

case "$3" in
    "0") printf "$BBLACK" ;;
    "1") printf "$BWHITE" ;;
    "2") printf "$BRED" ;;
    "3") printf "$BGREEN" ;;
    "4") printf "$BBLUE" ;;
    "5") printf "$BPURPLE" ;;
    "6") printf "$BBLACK" ;;
esac

printf " "

printf "\033[0m\n"
