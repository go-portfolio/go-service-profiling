package metrics

import "github.com/prometheus/client_golang/prometheus"

// RequestsTotal — счётчик общего числа HTTP-запросов к сервису
var (
	RequestsTotal prometheus.Counter
)

// Init инициализирует метрики Prometheus
func Init() {
	// Создаём новый счётчик
	RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",                   // имя метрики в Prometheus
		Help: "Total number of HTTP requests processed.", // описание метрики
	})

	// Регистрируем метрику в Prometheus
	// MustRegister — аварийно завершит программу, если регистрация не удалась
	prometheus.MustRegister(RequestsTotal)
}
