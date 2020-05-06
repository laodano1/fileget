package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
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

	// navigate to a page, wait for an element, click

	//var theBody string
	var res string
	var projects []*cdp.Node
	token := "aaaaaaaaaaaaaaaaaaaaaaa"
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://10.0.0.200:32567/login`),
		chromedp.Text(`button > span`, &res),
		//chromedp.Nodes(`#pane-sa > form > textarea`, &projects),

	//	// wait for footer element is visible (ie, page is loaded)
	//	//chromedp.WaitVisible(`body > #pane-sa > textarea`),
	//	// find and click "Expand All" link
	//	//chromedp.Click(`body > button`, chromedp.NodeVisible),
	//	// retrieve the value of the textarea




	//
	//	chromedp.Nodes(`body`, &projects),
	//
	)
	if err != nil {
		lg.Errorf("1 err: %v", err)
		return
	}
	lg.Debugf("res: %v", res)

	//if err := chromedp.Run(ctx, chromedp.Nodes(`textarea`, &projects)); err != nil {
	//)ByFunc(func(c context.Context, n *cdp.Node) (ids []cdp.NodeID, err error) {
	//	lg.Debugf("1 node name: %v", n.NodeType)
	//
	//	if strings.ToLower(n.NodeName) == "textarea" {
	//		lg.Debugf("node name: %v", n.NodeName)
	//		n.Value = token
	//		ids = append(ids, n.NodeID)
	//	}
	//	return
	//}))

	if err := chromedp.Run(ctx, chromedp.SetValue(`textarea`, token, chromedp.ByJSPath)); err != nil {
		if err != nil {
			lg.Errorf("2 err: %v", err)
			return
		}
	}
	//if err := chromedp.Run(ctx, chromedp.SetValue(`textarea`, token), ); err != nil {
	//	if err != nil {
	//		lg.Errorf("2 err: %v", err)
	//		return
	//	}
	//}
	//
	//lg.Debugf("111")
	//
	//if err := chromedp.Run(ctx, chromedp.Click(`button`, chromedp.NodeVisible)); err != nil {
	//	if err != nil {
	//		lg.Errorf("2 err: %v", err)
	//		return
	//	}
	//}

	lg.Debugf("button text: '%v', attrs: %v", res, projects[0].Attributes)
	for _, v := range projects {
		lg.Debugf("v: %v", v.NodeName)
		//lg.Debugf("v: %v", v.AttributeValue("placeholder"))
		//for _, attr := range v.Attributes {
		//	if attr == "placeholder" {
		//		lg.Debugf("placeholder: '%v'", v.AttributeValue(attr))
		//		break
		//	}
		//}
	}


}
