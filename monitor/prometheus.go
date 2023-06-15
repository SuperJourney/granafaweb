package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	histogramExample = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "myhistogram",
		Help:    "A histogram of request latencies in seconds.",
		Buckets: []float64{20, 40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 240, 260, 280, 300, 400, 500, 600, 700, 800, 900, 1000, 2000, 3000},
	}, []string{"path", "http_status"})

	summaryExample = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "mysummary",
		Help: "A summary of request latencies in seconds.",
		// MaxAge: 1 * time.Minute,
	}, []string{"path", "http_status"})

	counter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mycounter",
		Help: "A counter of request latencies in seconds.",
	}, []string{})

	gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mygauge",
		Help: "A gauge of request latencies in seconds.",
	}, []string{})
)

func Register() error {
	// Register a prometheus histogram
	if err := prometheus.Register(histogramExample); err != nil {
		return err
	}

	// Register a prometheus summary
	if err := prometheus.Register(summaryExample); err != nil {
		return err
	}

	// Register a prometheus counter
	if err := prometheus.Register(counter); err != nil {
		return err
	}

	// Register a prometheus gauge
	if err := prometheus.Register(gauge); err != nil {
		return err
	}
	return nil
}

func ObserveHistogramExample(value float64, lables ...string) {
	histogramExample.WithLabelValues(lables...).Observe(value)
}

func ObserveSummaryExample(value float64, lables ...string) {
	summaryExample.WithLabelValues(lables...).Observe(value)
}
