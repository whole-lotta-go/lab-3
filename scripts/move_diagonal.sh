#!/bin/bash

x=0.1
y=0.1

curl -X POST http://localhost:17000 -d "figure 0.1 0.1"

for i in {1..10}; do
    curl -X POST http://localhost:17000 -d "move $x $y"
    curl -X POST http://localhost:17000 -d "update"

    x=$(echo "$x + 0.05" | bc)
    y=$(echo "$y + 0.05" | bc)

    sleep 1
done
