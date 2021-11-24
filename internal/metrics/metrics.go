package metrics

import (
	"github.com/ozonmp/bss-office-facade/internal/config"
	"github.com/ozonmp/bss-office-facade/internal/model"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var totalEvents prometheus.Counter
var totalCud *prometheus.CounterVec

// InitMetrics - инициализирует метрики
func InitMetrics(cfg *config.Config) {
	totalEvents = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "processed_events_total",
		Help:      "Total number of processed events",
	})

	totalCud = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: cfg.Metrics.Namespace,
		Subsystem: cfg.Metrics.Subsystem,
		Name:      "processed_cud_total",
		Help:      "Total of events by type",
	}, []string{"type"})
}

// IncTotalEvents - увеличивает значение счетчика общего количества обработанных событий
func IncTotalEvents() {
	if totalEvents == nil {
		return
	}

	totalEvents.Inc()
}

// IncTotalCud - увеличивает значение счетчика событий конкретного типа
func IncTotalCud(eventType model.EventType) {
	if totalCud == nil {
		return
	}

	totalCud.WithLabelValues(string(eventType)).Inc()
}
