package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zsais/go-gin-prometheus"
)

func main() {
	// Create a new Prometheus registry
	reg := prometheus.NewRegistry()

	// Create a new histogram for request duration
	// with predefined buckets for milliseconds
	hist := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_milliseconds",
		Help:    "Duration of HTTP requests in milliseconds",
		Buckets: []float64{100, 200, 300, 400, 500, 750, 1000, 1500, 2000, 5000},
	}, []string{"method", "path"})

	// Register the histogram with the registry
	reg.MustRegister(hist)

	// Use the Prometheus middleware to instrument the router
	router := gin.Default()
	router.Use(ginprometheus.PromMiddleware(reg))

	// Setting up prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))

	// Starting the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
