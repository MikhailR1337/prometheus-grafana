package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "The total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
	ErrorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "The total number of HTTP errors",
		},
		[]string{"path", "method", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(ErrorCounter)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		method := c.Request.Method
		RequestCounter.WithLabelValues(path, method, http.StatusText(status)).Inc()
		if status >= http.StatusBadRequest {
			ErrorCounter.WithLabelValues(path, method, http.StatusText(status)).Inc()
		}
	}
}
