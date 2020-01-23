package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqCount = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "playground",
		Name:      "request_count",
		Help:      "Total number of request events",
	})
)

func main() {
	requestCount()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3030", nil)
}

func requestCount() {
	go func() {
		for {
			reqCount.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}
