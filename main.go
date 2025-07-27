package main

import (
	"errors"
	"net/http"
	"prometheus-grafan/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrInternal  = errors.New("internal server error")
	ErrForbidden = errors.New("forbidden")
)

func main() {
	router := gin.Default()

	middleware.PrometheusInit()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Use(middleware.PrometheusMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, from prometheus-grafana!"})
	})

	router.GET("/users", handler)

	router.GET("/comments", handler)

	router.GET("/posts", handler)

	router.Run(":8080")
}

func handler(c *gin.Context) {
	query := c.DefaultQuery("test", "")
	if query == "trigger-not-found" {
		c.JSON(http.StatusNotFound, gin.H{"error": ErrNotFound.Error()})
		return
	}
	if query == "trigger-forbidden" {
		c.JSON(http.StatusForbidden, gin.H{"error": ErrForbidden.Error()})
		return
	}
	if query == "trigger-server-error" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternal.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "everything is ok"})
}
