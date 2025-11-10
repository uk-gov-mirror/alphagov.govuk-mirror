package metrics

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

type Metrics struct {
	errorCounter prometheus.Counter
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		errorCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "crawler_errors_total",
			Help: "Total number of errors encountered by the crawler",
		}),
	}

	reg.MustRegister(m.errorCounter)

	return m
}

func UpdateErrorCounter(m *Metrics) {
	m.errorCounter.Inc()
}

func PushMetrics(reg prometheus.Registerer) {
	push.New(os.Getenv("PROMETHEUS_PUSHGATEWAY_URL"), "mirror_metrics").Push()
}
