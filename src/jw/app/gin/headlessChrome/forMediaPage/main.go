package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/davyxu/golog"
	"log"
	"time"
)

func main() {
	lg := golog.New("chromedp")
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://10.0.0.24:10000/`),
		chromedp.Click(`#btt`, chromedp.NodeVisible),
		//chromedp.Text(`button > span`, &res),
		//chromedp.SetValue(`textarea.el-textarea__inner`, token),
	)
	if err != nil {
		lg.Errorf("Run err: %v", err)
		return
	}
	lg.Debugf("res: %v", "1")


}
