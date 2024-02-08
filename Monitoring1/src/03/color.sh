#! /bin/bash

# Background
BBLACK="\033[40m"; BRED="\033[41m"; BGREEN="\033[42m"
BBLUE="\033[44m"; BPURPLE="\033[45m"; BWHITE="\033[47m"

# Foreground
TBLACK="\033[30m"; TRED="\033[31m"; TGREEN="\033[32m"
TBLUE="\033[34m"; TPURPLE="\033[35m"; TWHITE="\033[37m"

D="\033[0m"

TEMP_B1=$column1_background
TEMP_T1=$column1_font_color
TEMP_B2=$column2_background
TEMP_T2=$column2_font_color

while [ -n "$TEMP_B1" ]; do
  case "$TEMP_B1" in
    1) B1=$BWHITE ;;
    2) B1=$BRED ;;
    3) B1=$BGREEN ;;
    4) B1=$BBLUE ;;
    5) B1=$BPURPLE ;;
    6) B1=$BBLACK ;;
  esac
  break
done

while [ -n "$TEMP_T1" ]; do
  case "$TEMP_T1" in
    1) T1=$TWHITE ;;
    2) T1=$TRED ;;
    3) T1=$TGREEN ;;
    4) T1=$TBLUE ;;
    5) T1=$TPURPLE ;;
    6) T1=$TBLACK ;;
  esac
  break
done

while [ -n "$TEMP_B2" ]; do
  case "$TEMP_B2" in
    1) B2=$BWHITE ;;
    2) B2=$BRED ;;
    3) B2=$BGREEN ;;
    4) B2=$BBLUE ;;
    5) B2=$BPURPLE ;;
    6) B2=$BBLACK ;;
  esac
  break
done

while [ -n "$TEMP_T2" ]; do
  case "$TEMP_T2" in
    1) T2=$TWHITE ;;
    2) T2=$TRED ;;
    3) T2=$TGREEN ;;
    4) T2=$TBLUE ;;
    5) T2=$TPURPLE ;;
    6) T2=$TBLACK ;;
  esac
  break
done