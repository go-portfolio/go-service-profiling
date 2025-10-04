.PHONY: run build bench docker-build

run:
\tgo run ./cmd/server

build:
\tgo build -o bin/service ./cmd/server

bench:
\tgo test -bench=. -benchmem ./tests

docker-build:
\tdocker build -t go-service-profiling:latest .
