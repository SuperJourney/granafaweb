package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/SuperJourney/granafaweb/monitor"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := monitor.Register(); err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(monitor.RequestMonitor())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/hello", helloHandler)

	var f http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}

	router.GET("/demo", gin.WrapH(f))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func helloHandler(c *gin.Context) {
	num := rand.Intn(3000) + 1
	time.Sleep(time.Duration(num) * time.Millisecond)
	c.JSON(http.StatusOK, "hello,world")
}
