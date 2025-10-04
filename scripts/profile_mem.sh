#!/usr/bin/env bash
OUT=heap.prof
curl -s "http://localhost:8080/debug/pprof/heap" -o $OUT
echo "Saved $OUT. Run:"
echo "go tool pprof -http=:8081 ./bin/service $OUT"
