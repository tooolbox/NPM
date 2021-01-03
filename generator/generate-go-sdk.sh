#! /bin/bash

cd "$(dirname "$0")"

echo Delete the previously generated SDK code

rm -rf ../client
rm -rf ../models

echo Command to generate SDK

swagger generate client -A Fiserv -f ExternalYAML-6.14-v2-10-15-2020.yaml -t ../
