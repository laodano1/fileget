package main

import (
	"fileget/util"
	"flag"
	//"github.com/prometheus/client_golang"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	util.ZLogger.Debugf("metrics exposed...")
	util.ZLogger.Error(http.ListenAndServe(*addr, nil))
}
