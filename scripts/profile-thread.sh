#!/usr/bin/env bash
# ThreadCreate profile
curl -s "http://localhost:8080/debug/threadcreate" -o threadcreate.prof
echo "Saved threadcreate.prof"