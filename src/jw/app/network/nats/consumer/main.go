package main

import (
	"github.com/davyxu/golog"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

var (
	lg = golog.New("my.nats.consumer")
)

func main() {
	//nc, err := nats.Connect("10.0.0.146", nats.Name("API PublishBytes Example"))
	nc, err := nats.Connect("10.0.0.146")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	_, err = nc.Subscribe("updates", func(msg *nats.Msg) {
		lg.Infof("received msg: %v", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
