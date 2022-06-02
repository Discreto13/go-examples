package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	operationDone = promauto.NewCounter(prometheus.CounterOpts{
		Name: "promtestapp_operations_processed_total",
		Help: "Total number of processed operations",
	})
	waitPeriod = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "promtestapp_current_wait_period",
		Help: "Current period without answer",
	})
)

func startRecoringToMetrics() {
	for {
		var counter int
		rand.Seed(time.Now().UnixMilli())
		randPeriod := time.Duration(rand.Int()%10 + 2)
		slowTicker := time.After(time.Second * randPeriod) // randomly triggers between 3 and 8 seconds
		fastTicker := time.NewTicker(time.Second * 1)

		for {
			select {
			case t := <-slowTicker:
				operationDone.Inc()
				fmt.Printf("\n+1 operation done (%d:%d:%d)\n", t.Hour(), t.Minute(), t.Second())
				counter = 0
			case <-fastTicker.C:
				waitPeriod.Set(float64(counter))
				fmt.Printf("%d..", counter+1)
				counter++
			}
			if counter == 0 {
				break
			}
		}
	}
}

func main() {
	go startRecoringToMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
