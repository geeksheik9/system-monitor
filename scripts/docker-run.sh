#!/usr/bin/env bash
. ./scripts/version.sh

docker run -d -t -i -p 3001:3000 -e geeksheik9/login-service:$system_monitor_version