package prometheus

import (
	// 引入 prometheus 包
	"github.com/prometheus/client_golang/prometheus"
)

func Register() error {
	// Register a prometheus histogram
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "myhistogram",
		Help:    "A histogram of request latencies in seconds.",
		Buckets: []float64{.1, .2, .3, .4, .5, .6, .7, .8, .9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}, []string{})

	if err := prometheus.Register(histogram); err != nil {
		return err
	}

	// Register a prometheus summary
	summary := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "mysummary",
		Help: "A summary of request latencies in seconds.",
	}, []string{})
	if err := prometheus.Register(summary); err != nil {
		return err
	}

	// Register a prometheus counter
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mycounter",
		Help: "A counter of request latencies in seconds.",
	}, []string{})
	if err := prometheus.Register(counter); err != nil {
		return err
	}

	// Register a prometheus gauge
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mygauge",
		Help: "A gauge of request latencies in seconds.",
	}, []string{})
	if err := prometheus.Register(gauge); err != nil {
		return err
	}

	return nil
}
