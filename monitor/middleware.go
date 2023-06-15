package monitor

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		if c.Request.URL.Path == "/metrics" {
			return
		}

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		ObserveHistogramExample(float64(latency/time.Millisecond), c.Request.URL.Path, strconv.Itoa((c.Writer.Status())))
		ObserveSummaryExample(float64(latency.Microseconds()), c.Request.URL.Path, strconv.Itoa((c.Writer.Status())))

		fmt.Printf("Request Method: %s, Path: %s, Status: %d, Latency: %f\n",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), float64(latency)/1000000)
	}
}

// To use this middleware in a Gin router, you can add it like this:
// router.Use(monitor.RequestMonitor())
