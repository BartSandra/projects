#!/bin/bash

case $parametr in
 1) echo "$(cat ../04/access*.log 2>/dev/null | sort -k 9 -n)"
     ;;
 2) echo "$(cat ../04/access*.log 2>/dev/null | awk '!arr[$1]++' | awk '{print $1}')"
     ;;
 3) echo "$(cat ../04/access*.log 2>/dev/null | awk '$9~/^[4-5]/ {print $0}')"
     ;;
 4) echo "$(cat ../04/access*.log 2>/dev/null | awk '$9~/^[4-5]/ {print $0}' | awk '!arr[$1]++' | awk '{print $0}') "
     ;;
 *) echo "The parameter must be from 1 to 4"; exit 0;
esac
