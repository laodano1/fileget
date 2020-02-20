package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var (
	lg = golog.New("my.nats.publisher")
)

func main() {
	nc, err := nats.Connect("10.0.0.146", nats.Name("API PublishBytes Example"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		lg.Infof("drain all msg!")
		if err := nc.Drain(); err != nil {
			log.Fatal(err)
		}
		nc.Close()
	}()

	tc := time.Tick(3 * time.Second)
	for {
		select {
		case <-tc:
			pubInfo := fmt.Sprintf("msg | %v", time.Now().Format(time.RFC3339Nano))
			lg.Infof("publish info: %v", pubInfo)
			if err := nc.Publish("updates", []byte(pubInfo)); err != nil {
				log.Fatal(err)
			}
		}
	}
}
