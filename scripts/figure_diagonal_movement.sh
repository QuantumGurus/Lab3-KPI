#!/bin/bash

BASE_URL="http://localhost:17000"
curl -X POST "$BASE_URL" -d "reset"
curl -X POST "$BASE_URL" -d "white"

offset_x=$(echo "scale=2; 100/800" | bc)
offset_y=$(echo "scale=2; 100/800" | bc)

Xstart=$offset_x
Ystart=$offset_y

curl -X POST "$BASE_URL" -d "figure $Xstart $Ystart"

step=$(echo "scale=2; 20/800" | bc)

direction="down_right"

while true; do
    if [ "$direction" == "down_right" ]; then
        newX=$(echo "scale=2; $Xstart + $step" | bc)
        newY=$(echo "scale=2; $Ystart + $step" | bc)
        curl -X POST "$BASE_URL" -d "move $step $step"
        if (( $(echo "$newX > 1 - $offset_x" | bc) )) || (( $(echo "$newY > 1 - $offset_y" | bc) )); then
            direction="up_left"
            newX=$Xstart
            newY=$Ystart
            curl -X POST "$BASE_URL" -d "move $reverseStep $reverseStep"
        fi
    else
        newX=$(echo "scale=2; $Xstart - $step" | bc)
        newY=$(echo "scale=2; $Ystart - $step" | bc)
        reverseStep=$(echo "scale=2; $step * -1" | bc)
        curl -X POST "$BASE_URL" -d "move $reverseStep $reverseStep"
        if (( $(echo "$newX < $offset_x" | bc) )) || (( $(echo "$newY < $offset_y" | bc) )); then
            direction="down_right"
            newX=$Xstart
            newY=$Ystart
            curl -X POST "$BASE_URL" -d "move $step $step"
        fi
    fi

    Xstart=$newX
    Ystart=$newY
    curl -X POST "$BASE_URL" -d "update"

    sleep 0.1
done
