#!/usr/bin/env bash
# Скрипт для профилирования блокировок Go-сервиса через pprof

OUT=block.prof  # имя файла для сохранения профиля блокировок

# Запрос к эндпоинту block профиля
# -s — silent (без прогресса)
# -o $OUT — сохраняем результат в файл OUT
curl -s "http://localhost:8080/debug/pprof/block" -o $OUT

echo "Saved $OUT. Run:"

# Как открыть профайл через go tool pprof:
# -http=:8081 — открыть веб-интерфейс на порту 8081
# ./bin/service — путь к бинарнику Go-сервиса
# $OUT — файл с сохранённым профилем
echo "go tool pprof -http=:8081 ./bin/service $OUT"
