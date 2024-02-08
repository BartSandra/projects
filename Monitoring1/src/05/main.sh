#!/bin/bash

start="$(date +%s.%N)"

if [[ $# != 1 ]]; then
  echo "ERROR!!! Use 1 arguments!"
elif [[ $1 =~ /$ ]]; then
        source inf.sh $1
        end="$(date +%s.%N)"
        runtime=$(echo "$end - $start" | bc -l)
        echo "Script execution time (in seconds) = $runtime"
else
        echo "ERROR!!! Use /"
fi