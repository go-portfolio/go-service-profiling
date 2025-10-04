package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-service-profiling/internal/handlers"
	"go-service-profiling/internal/metrics"
)

func main() {
	// Инициализация метрик
	metrics.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/work", handlers.WorkHandler)
	mux.HandleFunc("/alloc", handlers.AllocHandler)
	mux.HandleFunc("/sleep", handlers.SleepHandler)

	// Метрики Prometheus
	mux.Handle("/metrics", promhttp.Handler())

	// pprof уже подключён через import _ "net/http/pprof"
	// доступно по /debug/pprof/

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
