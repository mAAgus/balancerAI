package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ai_balancer_requests_total",
			Help: "Total number of requests to AI balancer",
		},
		[]string{"type"},
	)
	RequestErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ai_balancer_requests_errors_total",
			Help: "Total number of failed requests",
		},
		[]string{"type"},
	)
	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ai_balancer_requests_duration_second",
			Help:    "Duration of request processing",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"type"},
	)
)
