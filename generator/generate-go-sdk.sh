#! /bin/bash

cd "$(dirname "$0")"

echo Delete the previously generated SDK code

rm -rf ../client
rm -rf ../models

echo Command to generate SDK

swagger generate client -A FirstVision -f swagger_20200820_508_0.yml -t ../
