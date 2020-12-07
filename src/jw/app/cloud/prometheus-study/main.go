package main

import (
	"fileget/util"
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"

	//"github.com/prometheus/client_golang"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	addr     = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	opsProc1 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "opsProc1_number",
		Help: "the number 1 for test",
	})

	opsProc2 = promauto.NewCounter(prometheus.CounterOpts{
		Name: "opsProc2_number",
		Help: "the number 2 for test",
	})

	//aa = prometheus.NewRegistry().
	tt = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "opsProc3_number",
		Help: "the number 3 for test",
	})
)

func recordMetrics() {
	go func() {
		for {
			util.ZLogger.Debugf("opsProc increased!")
			opsProc1.Inc()
			opsProc2.Add(2)
			//tt.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	flag.Parse()
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	util.ZLogger.Debugf("metrics exposed...")
	util.ZLogger.Error(http.ListenAndServe(*addr, nil))
}
