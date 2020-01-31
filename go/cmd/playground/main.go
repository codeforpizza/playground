package main

import (
	"flag"
	"fmt"
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
	// simple flag for port change
	var httpPort string
	flag.StringVar(&httpPort, "httpPort", ":3030", "http port for golang playground")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount()
		fmt.Fprint(w, "request recorded")
	})
	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Server running at :%s", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func requestCount() {
	go func() {
		for {
			reqCount.Inc()
			time.Sleep(1 * time.Second)
		}
	}()
}
