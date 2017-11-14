#! /bin/bash

for i in $(ps aux | grep "apache-geode" | awk '{print $2}' | head -n -1); do
        kill -9 $i
done