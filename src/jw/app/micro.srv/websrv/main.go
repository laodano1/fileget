package main

import (
	"github.com/micro/go-micro/web"
	"log"
)

func main() {

	mws := new(MyService)
	mws.myWebService = web.NewService(
		web.Name("go.micro.web.form"),
		func(o *web.Options) {
			o.Address = ":9999"
		},
	)

	mws.myWebService.Handle("/", WebServer.srv)
	if err := mws.myWebService.Run(); err != nil {
		log.Fatal("go micro web server error: ", err)
	}

}
