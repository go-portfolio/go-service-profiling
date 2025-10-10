#!/usr/bin/env bash
# Mutex profile
curl -s "http://localhost:8080/debug/mutex" -o mutex.prof
echo "Saved mutex.prof"
