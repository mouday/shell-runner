#!/bin/bash
 
# print time
for ((i=0; i<10; i++))
do
    echo $(date +"%Y-%m-%d %H:%M:%S")
    sleep 1
done