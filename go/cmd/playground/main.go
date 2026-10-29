package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// simple flag for port change
	var httpPort string
	flag.StringVar(&httpPort, "httpPort", ":3030", "http port for golang playground")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "request recorded")
	})

	fmt.Printf("Server running at :%s", httpPort)
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}
