# Go Profiling Service 🚀

Сервис для демонстрации профилирования CPU, памяти и задержек в Go, с интеграцией метрик Prometheus и поддержкой Docker.

---

## 🚀 Быстрый старт

### 1. Клонируйте репозиторий

```bash
git clone <your-repo-url>
cd go-service-profiling
```
2. Установите зависимости
```bash
go mod tidy
```
3. Запуск сервиса
```bash
make run
```
Сервис будет доступен на http://localhost:8080

## ⚙️ Доступные эндпоинты
| URL             | Описание                                                |
| --------------- | ------------------------------------------------------- |
| `/`             | Корневая страница, сообщает что сервис запущен          |
| `/work`         | Имитирует нагрузку CPU                                  |
| `/alloc`        | Имитирует выделение памяти и сборку мусора              |
| `/sleep`        | Имитирует задержку (sleep)                              |
| `/metrics`      | Экспорт метрик Prometheus                               |
| `/debug/pprof/` | Профилирование через pprof (CPU, heap, goroutine и др.) |


## 📊 Профилирование
CPU профайл
```bash
./scripts/profile-cpu.sh
# Сохранится cpu.prof
# Далее откроем веб-интерфейс:
go tool pprof -http=:8081 ./bin/service cpu.prof
```
Heap профайл (память)
```bash
./scripts/profile-heap.sh
# Сохранится heap.prof
# Веб-интерфейс:
go tool pprof -http=:8081 ./bin/service heap.prof
```
Веб-интерфейс pprof позволяет интерактивно изучать, какие функции потребляют CPU и память.

##  Бенчмарки
Запуск всех бенчмарков с измерением памяти:

bash
```
make bench
```
-bench=. — запуск всех бенчмарков

-benchmem — измерение использования памяти

## 🐳 Docker
Сборка Docker-образа:

```bash
make docker-build
```
Запуск контейнера:

```bash
docker run -p 8080:8080 go-service-profiling:latest
```
## 📈 Метрики Prometheus
Сервис экспортирует метрики на /metrics.
Пример метрики:

```nginx
http_requests_total 42
```
`http_requests_total` — общее количество HTTP-запросов к сервису


## Структура проекта
```bash
go-service-profiling/
├── cmd/server/          # main.go — точка входа
├── internal/
│   ├── handlers/        # HTTP-хендлеры
│   ├── metrics/         # метрики Prometheus
│   └── workload/        # функции имитации нагрузки (CPU, memory, sleep)
├── tests/               # тесты и бенчмарки
├── scripts/             # скрипты для профилирования (CPU, heap и др.)
├── bin/                 # собранный бинарник
├── Makefile             # сборка, запуск, бенчмарки, Docker
└── README.md
```
💡 Советы
- Используйте `/work`, `/alloc` и `/sleep` для генерации нагрузки перед профилированием.
- Открывайте профили через `go tool pprof -http=:8081`, чтобы видеть графики и горячие функции.
- Метрики `Prometheus` можно подключать к `Grafana` для визуализации.
