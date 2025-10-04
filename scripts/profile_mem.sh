#!/usr/bin/env bash
# Скрипт для сбора профиля памяти (heap) Go-сервиса через pprof

OUT=heap.prof  # имя файла, куда будет сохранён профиль памяти

# Запрос к эндпоинту pprof для heap-профиля
# -s — silent, чтобы не выводить лишний прогресс
# -o $OUT — сохраняем результат в файл OUT
curl -s "http://localhost:8080/debug/pprof/heap" -o $OUT

# Информация для пользователя
echo "Saved $OUT. Run:"

# Подсказка, как открыть сохранённый heap-профиль через go tool pprof
# -http=:8081 — открыть веб-интерфейс pprof на порту 8081
# ./bin/service — путь к бинарнику Go-сервиса
# $OUT — файл с сохранённым профилем памяти
echo "go tool pprof -http=:8081 ./bin/service $OUT"
