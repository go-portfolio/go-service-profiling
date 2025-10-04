package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"go-service-profiling/internal/workload"
	"go-service-profiling/internal/metrics"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Profiling Service is running 🚀")
	metrics.RequestsTotal.Inc()
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	workload.CPUHeavy(20000, 200)
	fmt.Fprintln(w, "done cpu work")
	metrics.RequestsTotal.Inc()
}

func AllocHandler(w http.ResponseWriter, r *http.Request) {
	workload.Allocate(1000, 1024*100) // 1000 объектов по 100KB
	runtime.GC()
	fmt.Fprintln(w, "allocated memory")
	metrics.RequestsTotal.Inc()
}

func SleepHandler(w http.ResponseWriter, r *http.Request) {
	d := time.Duration(100+rand.Intn(900)) * time.Millisecond
	workload.Sleep(d)
	fmt.Fprintln(w, "slept")
	metrics.RequestsTotal.Inc()
}
