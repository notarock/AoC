#!/usr/bin/env sh

red=12
green=13
blue=14

total=0

while read line; do
    picks=$(echo $line | sed "s/Game [0-9]*: \(.*\)/\1/g" | tr ",;" "\n" | awk '{$1=$1;print}')
    valid=true
    while IFS= read -r pick; do
        n=$(echo $pick | tr -dc '0-9')
        echo "Pick: $pick, n: $n"
        if [[ $pick == *"red"* ]]; then
            if [ $n -gt $red ]; then
                valid=false
                break
            fi
        elif [[ $pick == *"green"* ]]; then
            if [ $n -gt $green ]; then
                valid=false
                break
            fi
        elif [[ $pick == *"blue"* ]]; then
            if [ $n -gt $blue ]; then
                valid=false
                break
            fi
        fi
    done <<< "$picks"

    gameId=$(echo $line | sed "s/Game \([0-9]*\).*/\1/g")
    if [ $valid = "true" ]; then
        total=$((total + $gameId))
        echo "Game $gameId: valid"
    else
        echo "Game $gameId: invalid"
    fi

done <input.txt

echo "Total = $total"
