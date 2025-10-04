package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestsTotal prometheus.Counter
)

func Init() {
	RequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests processed.",
	})
	prometheus.MustRegister(RequestsTotal)
}
