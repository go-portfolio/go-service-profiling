# Файл Makefile для управления сборкой и запуском Go-сервиса

# .PHONY — это "фиктивные" цели, которые не соответствуют файлам
# Позволяет Make всегда выполнять эти команды, даже если есть файлы с такими именами
.PHONY: run build bench docker-build

# run — запуск сервиса напрямую через go run
run:
	go run ./cmd/server
	# ./cmd/server — путь к точке входа сервиса (main.go)

# build — компиляция сервиса в бинарник
build:
	go build -o bin/service ./cmd/server
	# -o bin/service — результат компиляции будет в bin/service

# bench — запуск всех бенчмарков в папке tests
bench:
	go test -bench=. -benchmem ./tests
	# -bench=. — запуск всех функций бенчмарков
	# -benchmem — измерять использование памяти при выполнении бенчмарков

# docker-build — сборка Docker-образа сервиса
docker-build:
	docker build -t go-service-profiling:latest .
	# -t go-service-profiling:latest — тег образа
	# . — контекст сборки (текущая папка)
