package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // подключение стандартного профилировщика pprof

	"github.com/go-portfolio/go-service-profiling/internal/handlers" // пакет с HTTP-хендлерами
	"github.com/go-portfolio/go-service-profiling/internal/metrics"  // пакет с метриками
	"github.com/prometheus/client_golang/prometheus/promhttp"        // адаптер для экспорта метрик в Prometheus
)

func main() {
	// Инициализация пользовательских метрик (счётчики, гистограммы и т.д.)
	metrics.Init()

	// Создаём новый маршрутизатор HTTP
	mux := http.NewServeMux()

	// Регистрируем обработчики (эндпоинты сервиса)
	mux.HandleFunc("/", handlers.IndexHandler)   // корневая страница
	mux.HandleFunc("/work", handlers.WorkHandler) // имитация нагрузки
	mux.HandleFunc("/alloc", handlers.AllocHandler) // выделение памяти (тест профилировки)
	mux.HandleFunc("/sleep", handlers.SleepHandler) // "усыпление" горутины (тест задержек)

	// Подключаем метрики Prometheus по пути /metrics
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	
	// Профилировщик pprof уже подключён выше через blank import,
	// доступен по адресу /debug/pprof/
	// Пример: http://localhost:8080/debug/pprof/

	// Запускаем HTTP-сервер на порту 8080
	addr := ":8080"
	log.Printf("listening on %s", addr)

	// Если сервер завершится с ошибкой — логируем фатально и выходим
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
