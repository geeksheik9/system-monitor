#!/usr/bin/env bash

cat /sys/class/thermal/thermal_zone*/temp | column -s $'\t' -t | sed 's/\(.\)..$/.\1/'