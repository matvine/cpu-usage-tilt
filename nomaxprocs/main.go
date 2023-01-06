package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/cpu.eater", cpuEaterHandler)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8009", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	cores := runtime.NumCPU()
	fmt.Fprintf(w, "Number of Cores = %d", cores)

}

func cpuEaterHandler(w http.ResponseWriter, r *http.Request) {

	started := time.Now()
	done := make(chan int, 1)

	go func() {
		for i := 0; i < runtime.NumCPU(); i++ {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}
	}()

	time.Sleep(time.Millisecond * 250)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request duration %v", time.Since(started))
}
