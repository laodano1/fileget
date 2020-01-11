package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	application struct {
		Name string
		sdcli   *api.Client // service discovery client
		sdsrv   *echo.Echo  // for sd http server
		srv     *echo.Echo    // server echo http server
	}
)

var (
	//consulHost = "192.168.11.67:8500/"
	consulHost = "10.0.0.34:8500"
)

var logger = golog.New("service.discovery")

func newDefaultApp(name string, consul string) (app *application) {
	app = new(application)
	config := api.DefaultConfig()
	config.Address = consul

	logger.Infof("New consul client")
	client, err := api.NewClient(config)
	if err != nil {
		logger.Errorf("new client error: %v", err)
	}

	app.Name = name
	app.sdcli = client
	logger.Infof("New client agent")
	//agent := client.Agent()

	return
}

func (a *application) SrvRigister(srvName string, port string) {
	//注册到consul的服务器内容
	logger.Infof("fill in register contents")
	portNum, _ := strconv.Atoi(port)

	meta := make(map[string]string)
	meta["cd"] = "test1"
	meta["nk"] = "test2"
	reg := &api.AgentServiceRegistration{
		ID:   "jw-srv-10.0.0.31:" + port,
		Name: srvName,
		Port: portNum,
		Address: "10.0.0.31",
		Meta:  meta,
		Tags: []string{"gggaaammmeee", "网关"},
		Check: &api.AgentServiceCheck{
			CheckID:  port,
			Name:     "srv-chk-name",
			TTL:                            "3s",
			DeregisterCriticalServiceAfter: "60s",
		},
	}

	go func() {
		tk := time.Tick(2 * time.Second)
		//td := time.Tick(120 * time.Second)
		cnt := 0
		for {
			select {
			case <- tk:
				cnt++
				a.sdcli.Agent().PassTTL(reg.Check.CheckID, fmt.Sprintf("hello -> %d", cnt))
			//case  <- td:
			//
			//	return
			}
		}
	}()

	logger.Infof("register to consul")
	//注册到consul
	if err := a.sdcli.Agent().ServiceRegister(reg); err != nil {
		log.Fatal("register failure: ", err)
	}
}

func (a *application) srvStart() {
	go startSRV()
}


func listHandler(response http.ResponseWriter, request *http.Request)  {
	response.Write([]byte("Hello hello!"))
}

func startSRV() {
	logger.Infof("new http server with  2")
	http.HandleFunc("/srv", listHandler)

	if tmpErr := http.ListenAndServe(":8888", nil); tmpErr != nil {
		panic(fmt.Sprintf("callback server Error:%v", tmpErr))
	}
}

func main1() {
	app := newDefaultApp("jw-app", consulHost)   // app name

	done := make(chan bool)
	go func() {
		tk := time.Tick(30 * time.Second)
		regPort := 8888
		for {
			select {
			case <- tk:
				app.SrvRigister("wangke-agent", fmt.Sprintf("%d", regPort))  // service register
				regPort++
			case <- done:
				return
			}
		}
	}()

	//wg := sync.WaitGroup{}
	app.srvStart()
	//wg.Wait()

	//app.sdcli.Agent().


	time.Sleep(3 * time.Minute)
	//time.AfterFunc(3 * time.Minute, func() {
		done <- true
	//})
	log.Println("Bye bye!!!")
	//
	////app.sd.Health().Service()

}
