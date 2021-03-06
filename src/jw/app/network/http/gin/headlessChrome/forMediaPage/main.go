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
	ctx, cancel = context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	var res string
	var class string
	//var ok    bool

	err := chromedp.Run(ctx,
			//chromedp.Navigate(`https://www.cnblogs.com/dhcn/`),
		chromedp.Navigate(`http://10.0.0.200:32567/namespace/wanda-test-use`),
		chromedp.WaitVisible(`html`),
		//chromedp.AttributeValue(`body`, "style", &class, &ok),
		//chromedp.Navigate(`https://micro.mu/docs/go-grpc.html`),
		//chromedp.WaitVisible(`#overview`),
		chromedp.Title(&res),


		//chromedp.SetValue(`textarea.el-textarea__inner`, token),
	)
	if err != nil {
		lg.Errorf("1-Run err: %v", err)
		return
	}
	lg.Debugf("res: %v", res)
	lg.Debugf("class: %v", class)

	//if err = chromedp.Run(ctx,
	//	chromedp.Navigate(`https://www.autohome.com.cn/car/`),
	//	//chromedp.WaitVisible(`#overview`),
	//
	//	//
	//	); err != nil {
	//	lg.Errorf("2-Run err: %v", err)
	//	return
	//}


	//if err := chromedp.Run(ctx, chromedp.Tasks{
	//
	//	chromedp.ActionFunc(func(c context.Context) (err error) {
	//		lg.Debugf("Navigate")
	//		chromedp.Navigate(`https://www.autohome.com.cn/car/#pvareaid=3311275`)
	//		//lg.Debugf("Click")
	//		//chromedp.Click(`#blog_nav_myhome`, chromedp.NodeVisible)
	//		//lg.Debugf("WaitVisible")
	//		//chromedp.WaitVisible(`#navigator`)
	//		lg.Debugf("Text")
	//		chromedp.Text(`#price-range`, &res)
	//		return
	//	}),
	//
	//}); err != nil {
	//	lg.Errorf("run Tasks err: %v", err)
	//	return
	//}

	lg.Debugf("res: %v", res)

}
