package main

import (
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"html/template"
	"net/http"
	"time"
)

const (
	srvAdd = ":9999"
)

var logger = golog.New("my-web-server")

func newWebSrv() *webSrvObj {
	return &webSrvObj{
		name:   "my-gin",
		e:      gin.Default(),
	}
}

func (s *webSrvObj)  hp(c *gin.Context) {
	tmpl := template.Must(template.New("aaa").Parse(hptemplate))

	var sts []gin.H

	if s.cor.acs.cliNum <= 0 {

	} else {
		//for i, v := range s.cor.acs.asMap {
		for i := 0; i < s.cor.acs.cliNum; i++{
			st := gin.H{
				"cid"   : i,
				"log": "this is log",
				"phase" : s.cor.acs.asMap[i].phase,
			}
			sts = append(sts, st)
		}
	}

	dt := gin.H{
		"title": "hello world",
		"log": "this is log",
		"clientNum": s.cor.acs.cliNum,
		"allStatus": sts,
	}

	//
	rd := render.HTML{
		Template: tmpl,
		Name:     "aaa",
		Data:     dt,
	}
	c.Render(http.StatusOK, rd)

	//c.IndentedJSON(http.StatusOK, sts)

}

//func (s *webSrvObj)  allClientStatus() {
//
//
//}

//func (s *webSrvObj)  getAllClientStatus(ctx *gin.Context) {
//
//	ctx.String(http.StatusOK, fmt.Sprintf("%v", s.))
//}

func (cho *chObj) startClients() {

	for i := 0; i < 10; i++ {
		go func(id int) {
			//rand.Seed(time.Now().UnixNano())
			//for rand.Intn(6)
			itv := time.Now().Nanosecond() % 8
			if itv < 4 {
				itv += 3
			}
			logger.Debugf("random interval: %v", itv)
			tk := time.Tick(time.Duration(itv) * time.Second)
			ph := map[int]string{
				0: "开",
				1: "叫",
				2: "抢",
				3: "出",
			}
			p := 0
			for {
				select {
				case <- tk:
					p %= 4
					cho.stch <- &clientStatus{
						id: id,
						phase: ph[p],
					}
					p++
				}
			}
		}(i)
	}

	//dt := `{"status":"OK"}`
	//c.JSON(http.StatusOK, dt)
}

func (cho *chObj) dealCliStatus()  {
	cho.acs.asMap = make(map[int]*clientStatus)
	for {
		select {
		case item := <- cho.stch:
			if  _, ok := cho.acs.asMap[item.id]; !ok {
				cho.acs.cliNum++
			}
			cho.acs.asMap[item.id] = item

		}
	}
}

func newChObj() *chObj {
	return &chObj{
		stch: make(chan *clientStatus),
	}
}


func main() {

	//logger.SetParts()
	gin.ForceConsoleColor()
	mws := newWebSrv()   // web server
	co  := newChObj()    // message queue
	co.startClients()    // start workers
	go co.dealCliStatus()
	mws.cor = co

	mws.e.GET("/", mws.hp)

	mws.e.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	if err := mws.e.Run(srvAdd); err != nil {
		logger.Errorf("router run failed: %v", err)
	}
}