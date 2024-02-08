#! /bin/bash

#source inf.sh
#source color.sh

check_parametr=$(bash check.sh)

color=$(bash color.sh ${check_parametr[@]:0:7})

bash inf.sh $color

bash inf_check.sh $check_parametr

