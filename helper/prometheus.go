package helper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ActiveMonitors a counter with the current loaded monitors
	ActiveMonitors prometheus.Counter

	// DegradedMonitors a gauge with the current degraded monitors
	DegradedMonitors prometheus.Gauge
)

func init() {
	ActiveMonitors = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "active_monitors",
			Help: "provides a counter with the current loaded monitors",
		},
	)

	DegradedMonitors = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "degraded_monitors",
			Help: "provides a gauge with the current degraded monitors",
		},
	)
}
