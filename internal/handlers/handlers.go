package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/go-portfolio/go-service-profiling/internal/metrics"  // метрики Prometheus
	"github.com/go-portfolio/go-service-profiling/internal/workload" // функции имитации нагрузки
)

// IndexHandler — корневой обработчик, возвращает простое сообщение о работе сервиса
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Profiling Service is running 🚀") // выводим сообщение в ответ
	metrics.RequestsTotal.Inc()                            // увеличиваем общий счётчик запросов
}

// WorkHandler — имитация CPU-нагрузки
func WorkHandler(w http.ResponseWriter, r *http.Request) {
	workload.CPUHeavy(20000, 200) // выполняем тяжёлые вычисления (20000 итераций по 200 циклов)
	fmt.Fprintln(w, "done cpu work") // сообщаем, что работа выполнена
	metrics.RequestsTotal.Inc()      // увеличиваем счётчик запросов
}

// AllocHandler — имитация выделения памяти
func AllocHandler(w http.ResponseWriter, r *http.Request) {
	workload.Allocate(1000, 1024*100) // выделяем 1000 объектов по 100KB каждый
	runtime.GC()                       // вызываем сборку мусора для очистки памяти
	fmt.Fprintln(w, "allocated memory") // сообщаем, что память выделена
	metrics.RequestsTotal.Inc()         // увеличиваем счётчик запросов
}

// SleepHandler — имитация задержки/ожидания
func SleepHandler(w http.ResponseWriter, r *http.Request) {
	// случайная задержка от 100 до 1000 миллисекунд
	d := time.Duration(100+rand.Intn(900)) * time.Millisecond
	workload.Sleep(d)                   // "усыпляем" горутину на указанное время
	fmt.Fprintln(w, "slept")            // сообщаем, что задержка завершена
	metrics.RequestsTotal.Inc()         // увеличиваем счётчик запросов
}
