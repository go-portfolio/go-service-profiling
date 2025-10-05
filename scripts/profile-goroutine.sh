#!/usr/bin/env bash
# Скрипт для сбора профиля горутин Go-сервиса через pprof

OUT=goroutine.prof  # имя файла для сохранённого профиля

# Запрос к эндпоинту pprof для получения состояния горутин
# -s — silent режим (без прогресса)
# -o $OUT — сохранить результат в файл
curl -s "http://localhost:8080/debug/pprof/goroutine?debug=2" -o $OUT

echo "Saved $OUT. Run:"
echo "go tool pprof -http=:8081 ./bin/service $OUT"
