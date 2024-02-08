#!/bin/bash

TEMP_B1=$(grep column1_background color.conf | cut -b 20-)
TEMP_T1=$(grep column1_font_color color.conf | cut -b 20-)
TEMP_B2=$(grep column2_background color.conf | cut -b 20-)
TEMP_T2=$(grep column2_font_color color.conf | cut -b 20-)

if [[ "$TEMP_B1" == "$TEMP_T1" || "$TEMP_B2" == "$TEMP_T2" ]]; then
        echo "ERROR!!!"
    else
    if ! [[ $TEMP_B1 =~ [1-6] ]]; then
        TEMP_B1=0
    fi

    if ! [[ $TEMP_T1 =~ [1-6] ]]; then
        TEMP_T1=0
    fi

    if ! [[ $TEMP_B2 =~ [1-6] ]]; then
        TEMP_B2=0
    fi

    if ! [[ $TEMP_T2 =~ [1-6] ]]; then
        TEMP_T2=0
    fi
        echo $TEMP_B1 $TEMP_T1 $TEMP_B2 $TEMP_T2
    fi