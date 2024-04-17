#!/bin/bash

#chmod u+x download.sh

mkdir -p /data/hist/
mc mirror ny2/edi/e2020/ ./data/ --overwrite 

# Directory to search for zip files
directory="./data/hist/lst_evt"
cd "$directory" || { echo "Failed to change directory to $directory"; exit 1; }

for zipFile in *.zip; do
    if [ -f "$zipFile" ]; then
        unzip -q -o "$zipFile"

        if [ $? -eq 0 ]; then
            rm "$zipFile"
        else
            echo "Error unzipping file: $zipFile"
        fi
    else
        echo "No .zip files found in $directory"
        break
    fi
done