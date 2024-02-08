#!/bin/bash

# 1-Все записи, отсортированные по коду ответа 
# 2-Все уникальные IP, встречающиеся в записях
# 3-Все запросы с ошибками (код ответа - 4хх или 5хх) 
# 4-Все уникальные IP, которые встречаются среди ошибочных запросов

if [ $# -ne 1 ]
 then
  echo "The script is run with one parameter!" >&2
  exit 0
fi

export parametr=$1

bash monitoring.sh $1