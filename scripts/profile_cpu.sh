#!/usr/bin/env bash
# Скрипт для профилирования CPU Go-сервиса через pprof

OUT=cpu.prof  # имя файла, куда будет сохранён профиль CPU

# Запрос к эндпоинту pprof профилирования CPU на 30 секунд
# -s — silent, чтобы не выводить прогресс
# -o $OUT — сохраняем результат в файл OUT
curl -s "http://localhost:8080/debug/pprof/profile?seconds=30" -o $OUT

# Информация для пользователя
echo "Saved $OUT. Run:"

# Подсказка, как открыть профайл через go tool pprof
# -http=:8081 — открыть веб-интерфейс на порту 8081
# ./bin/service — путь к бинарнику Go-сервиса
# $OUT — файл с сохранённым профилем
echo "go tool pprof -http=:8081 ./bin/service $OUT"
