#!/usr/bin/env sh

red=12
green=13
blue=14

total=0
pTotal=0

while read line; do
    picks=$(echo $line | sed "s/Game [0-9]*: \(.*\)/\1/g" | tr ",;" "\n" | awk '{$1=$1;print}')
    valid=true

    minRed=0
    minGreen=0
    minBlue=0

    while IFS= read -r pick; do
        n=$(echo $pick | tr -dc '0-9')
        if [[ $pick == *"red"* ]]; then
            if [ $n -gt $minRed ]; then
                minRed=$n
            fi
            if [ $n -gt $red ]; then
                valid=false
            fi
        elif [[ $pick == *"green"* ]]; then
            if [ $n -gt $minGreen ]; then
                minGreen=$n
            fi
            if [ $n -gt $green ]; then
                valid=false
            fi
        elif [[ $pick == *"blue"* ]]; then
            if [ $n -gt $minBlue ]; then
                minBlue=$n
            fi
            if [ $n -gt $blue ]; then
                valid=false
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

    echo "Min red = $minRed, min green = $minGreen, min blue = $minBlue"

    power=$((minRed * minGreen * minBlue))
    pTotal=$((pTotal + $power))
    echo "Game $gameId: power = $power"

done <input.txt

echo "Total = $total"
echo "pTotal = $pTotal"
