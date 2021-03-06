package main

import (
	"flag"
	"github.com/davyxu/golog"
	"github.com/micro/go-micro/web"
	"log"
)

var (
	PORT = "0.0.0.0:9999"
	lg   = golog.New("go.micro.web.srv")
)

func main() {
	var mode string
	flag.StringVar(&mode, "m", "web", "web server work mode. default mode is http. other mode can support websocket")
	flag.Parse()

	lg.Infof("mode: %v", mode)

	mws := new(MyService)
	mws.myWebService = web.NewService(
		web.Name("go.micro.web.form"),
		web.Address(PORT),
		//func(o *web.Options) {
		//	o.Address = PORT
		//},
	)

	mws.myWebService.Handle("/", WebServer.srv)
	if err := mws.myWebService.Run(); err != nil {
		log.Fatal("go micro web server error: ", err)
	}

}
