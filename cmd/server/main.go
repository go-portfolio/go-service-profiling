package main

import (
	"log"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/go-portfolio/go-service-profiling/internal/handlers"
	"github.com/go-portfolio/go-service-profiling/internal/metrics"
	"github.com/go-portfolio/go-service-profiling/profiling"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Инициализация пользовательских метрик
	metrics.Init()

	// Создаём HTTP маршрутизатор
	mux := http.NewServeMux()

	// Основные эндпоинты сервиса
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/work", handlers.WorkHandler)
	mux.HandleFunc("/alloc", handlers.AllocHandler)
	mux.HandleFunc("/sleep", handlers.SleepHandler)

	// Prometheus метрики
	mux.Handle("/metrics", promhttp.Handler())

	// Pprof эндпоинты
	mux.Handle("/debug/pprof/", http.DefaultServeMux)

	// -----------------------------
	// 🔹 Постоянное профилирование
	// -----------------------------

	// Block profile (для /debug/pprof/block)
	profiling.EnableBlockProfile(1)

	// Периодическое снятие heap профиля
	go func() {
		for {
			time.Sleep(60 * time.Second)
			if err := profiling.WriteHeapProfile("heap.prof"); err != nil {
				log.Println("heap profile error:", err)
			} else {
				log.Println("heap profile saved to heap.prof")
			}
		}
	}()

	// -----------------------------
	// 🔹 Эндпоинты для ручного запуска CPU и Trace
	// -----------------------------

	// CPU profile на 30 секунд
	mux.HandleFunc("/debug/cpu", func(w http.ResponseWriter, r *http.Request) {
		f, err := profiling.StartCPUProfile("cpu.prof")
		if err != nil {
			http.Error(w, "failed to start CPU profile", http.StatusInternalServerError)
			return
		}
		defer profiling.StopCPUProfile(f)

		w.Write([]byte("CPU profiling started for 30s...\n"))
		time.Sleep(30 * time.Second)
		w.Write([]byte("CPU profiling finished, saved to cpu.prof\n"))
	})

	// Trace на 5 секунд
	mux.HandleFunc("/debug/trace", func(w http.ResponseWriter, r *http.Request) {
		tf, err := profiling.StartTrace("trace.out")
		if err != nil {
			http.Error(w, "failed to start trace", http.StatusInternalServerError)
			return
		}
		defer profiling.StopTrace(tf)

		w.Write([]byte("Trace started for 5s...\n"))
		time.Sleep(5 * time.Second)
		w.Write([]byte("Trace finished, saved to trace.out\n"))
	})

	// -----------------------------
	// Запуск сервера
	// -----------------------------
	addr := ":8080"
	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
