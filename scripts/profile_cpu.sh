#!/usr/bin/env bash
OUT=cpu.prof
curl -s "http://localhost:8080/debug/pprof/profile?seconds=30" -o $OUT
echo "Saved $OUT. Run:"
echo "go tool pprof -http=:8081 ./bin/service $OUT"
