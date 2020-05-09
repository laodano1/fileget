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

	// navigate to a page, wait for an element, click
	//var projects []*cdp.Node
	var msg string
	var res string
	var target []string
	//ok := true
	//token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrdWJvYXJkLXVzZXItdG9rZW4tZndsc2QiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoia3Vib2FyZC11c2VyIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNzdjZTE5NmYtMjljYy00NWExLWI4NGQtNzM1Zjc0NGM0ZDJmIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmt1Ym9hcmQtdXNlciJ9.k3A7nD3KRKVzBuMMJEKRzUE2PUUA3-lkJgB1-xip2qsto-Ohjx_XlAjjeNcaopD3NgVrTBGh2PpF6gGrdr_OEYhcpFzja_E-frWyIKnH8UcIWs486LMECM-BosXd1sqhsBQphTWQf-5EDdWDMMQbzlQko5yWGqbHwm5SXPlXt3waKZnRGGijB8JeQ0wiVj-aP2FHLclxoc7hMDp1vVoBSljYlCjzpGxewgrweqhBQYjUELTgvTlAGRXzC0qOj25yLrF8v1EOMkNBvYNGzPI1QFmsiv2OmnnTGGLRoM8cqPLusZyfjsdDXvZDL0E2bVL7MSi9e2q_7QLBn6kfWmDjXw"
	token := "aaaaaaaaaaaaaaa"
	//var innerhtml string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://10.0.0.200:32567/login`),
		chromedp.Text(`#tab-sa`, &res),

	    chromedp.SetValue(`textarea.el-textarea__inner`, token, chromedp.NodeVisible),
		//chromedp.Click(`button.el-button.el-button--primary.el-button--mini`, chromedp.NodeVisible),
		chromedp.Evaluate(`Object.keys(window);`, &target),
		//chromedp.AttributeValue(`a.link`, "target", &target, &ok),
		//chromedp.AttributeValue(`#__BVID__94`, "class", &target, &ok),

		//chromedp.ActionFunc(func(c context.Context) (err error) {
		//	time.Sleep(2 * time.Second)
		//	return
		//}),
		//chromedp.Text(`div.el-form-item__error`, &msg),
	//
	)
	if err != nil {
		lg.Errorf("1 err: %v", err)
		return
	}
	lg.Debugf("res: %v", target)

	lg.Debugf("item__error: %v", msg)
	//if err = chromedp.Run(ctx, chromedp.Title(&target),); err != nil {
	//	if err != nil {
	//		lg.Errorf("target err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("target: %v", target)




	//var placeholder string
	//if err = chromedp.Run(ctx, chromedp.AttributeValue(`textarea.el-textarea__inner`, "placeholder", &placeholder, &ok)); err != nil {
	//	if err != nil {
	//		lg.Errorf("placeholder err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("placeholder: %v", placeholder)
	//
	//if err = chromedp.Run(ctx, chromedp.SetValue(`textarea.el-textarea__inner`, token)); err != nil {
	//	if err != nil {
	//		lg.Errorf("placeholder err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("placeholder: %v", msg)

	//if err := chromedp.Run(ctx, chromedp.Value(`textarea.el-textarea__inner`, &msg)); err != nil {
	//	if err != nil {
	//		lg.Errorf("placeholder err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("Value: %v", msg)



	//if err = chromedp.Run(ctx, chromedp.AttributeValue(`textarea.el-textarea__inner`, "placeholder", &msg, &ok)); err != nil {
	//if err = chromedp.Run(ctx, chromedp.Value(`textarea.el-textarea__inner`, &msg)); err != nil {
	//	if err != nil {
	//		lg.Errorf("placeholder err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("token: %v", msg)

	//var typeValue string ="null"
	//if err = chromedp.Run(ctx, chromedp.Click(`#pane-sa > div > form > div > button.el-button.el-button--primary.el-button--mini`, chromedp.NodeVisible)); err != nil {
	//if err = chromedp.Run(ctx, chromedp.Click(`button.el-button.el-button--primary.el-button--mini`, chromedp.NodeVisible)); err != nil {
	//	if err != nil {
	//		lg.Errorf("Click err: %v", err)
	//		return
	//	}
	//}
	//lg.Debugf("type: %v", "---")

	//var val string
	if err = chromedp.Run(ctx, chromedp.Tasks{
		//chromedp.ActionFunc(func(c context.Context) (err error) {
		//	time.Sleep(2 * time.Second)
		//	return
		//}),
		//chromedp.Value(`textarea.el-textarea__inner`, &val),

		//chromedp.Click(`button.el-button.el-button--primary.el-button--mini`, chromedp.NodeVisible),

		//chromedp.ActionFunc(func(c context.Context) (err error) {
		//	time.Sleep(2 * time.Second)
		//	return
		//}),

		//chromedp.Text(`div.el-form-item__error`, &msg),

	}); err != nil {
		lg.Errorf("run Tasks err: %v", err)
		return
	}

	//lg.Debugf("token: %v", val)



}
