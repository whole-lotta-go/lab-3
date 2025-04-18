#!/bin/bash

x=0.05
y=0.05
dx=0.005
dy=0.005

curl -X POST http://localhost:17000 -d "figure 0 0"

while true; do
    curl -X POST http://localhost:17000 -d "move $x $y"
    curl -X POST http://localhost:17000 -d "update"

    x=$(echo "$x + $dx" | bc)
    y=$(echo "$y + $dy" | bc)


    if (( $(echo "$x >= 0.8" | bc -l) )); then
        x=0.8
        dx=$(echo "-($dx)" | bc -l)
    elif (( $(echo "$x <= 0.2" | bc -l) )); then
        x=0.2
        dx=$(echo "-($dx)" | bc -l)
    fi

    if (( $(echo "$y >= 0.8" | bc -l) )); then
        y=0.8
        dy=$(echo "-($dy)" | bc -l)
    elif (( $(echo "$y <= 0.2" | bc -l) )); then
        y=0.2
        dy=$(echo "-($dy)" | bc -l)
    fi
done
