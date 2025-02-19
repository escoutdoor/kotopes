package metrics

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "kotopes"
	appName   = "auth"
)

type Metrics struct {
	requestCounter    prometheus.Counter
	responseCounter   *prometheus.CounterVec
	histogramRespTime *prometheus.HistogramVec
}

var metrics *Metrics

func Init(_ context.Context) error {
	metrics = &Metrics{
		requestCounter: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "grpc",
			Name:      appName + "_requests_total",
			Help:      "Number of requests to the server",
		}),
		responseCounter: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "grpc",
			Name:      appName + "_responses_total",
			Help:      "Number of responses from the server",
		}, []string{"status", "method"}),
		histogramRespTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "grpc",
			Name:      appName + "_histogram_response_time_seconds",
			Help:      "Server time response",
			Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
		}, []string{"status"}),
	}
	return nil
}

func IncRequestCounter() {
	metrics.requestCounter.Inc()
}

func IncResponseCounter(status string, method string) {
	metrics.responseCounter.WithLabelValues(status, method).Inc()
}

func HistogramResponseTimeObserve(status string, time float64) {
	metrics.histogramRespTime.WithLabelValues(status).Observe(time)
}
