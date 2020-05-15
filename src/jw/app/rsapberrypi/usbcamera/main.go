package main

import (
	"fmt"
	"github.com/blackjack/webcam"
	"os"
)

func main() {
	cam, err := webcam.Open("/dev/video0") // Open webcam
	if err != nil {
		panic(err.Error())
	}
	defer cam.Close()
	// ...
	// Setup webcam image format and frame size here (see examples or documentation)
	// ...
	err = cam.StartStreaming()
	if err != nil {
		panic(err.Error())
	}

	timeout := uint32(100)
	for {
		err = cam.WaitForFrame(timeout)
		switch err.(type) {
		case nil:
		case *webcam.Timeout:
			fmt.Fprint(os.Stderr, err.Error())
			continue
		default:
			panic(err.Error())
		}

		frame, err := cam.ReadFrame()
		if len(frame) != 0 {
			// Process frame

		} else if err != nil {
			panic(err.Error())
		}
	}
}
