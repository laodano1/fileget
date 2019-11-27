package main

import (
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
)

type (

	application struct {
		Name string
		sd *api.Client
		srv *echo.Echo
	}

)

var (
	consulHost = "192.168.1.146:8500"
)

var logger = golog.New("service.discovery")

func newDefaultApp(name string) (app *application) {
	app = new(application)
	config := api.DefaultConfig()
	config.Address = consulHost

	logger.Infof("New consul client")
	client, err := api.NewClient(config)
	if err != nil {
		logger.Errorf("new client error: %v", err)
	}

	app.Name = name
	app.sd = client
	logger.Infof("New client agent")
	//agent := client.Agent()


	return
}

func (a *application) SrvRigister(srvName string, port string) {
	//注册到consul的服务器内容
	logger.Infof("fill in register contents")
	portNum, _ :=  strconv.Atoi(port)
	reg := &api.AgentServiceRegistration{
		ID: "jinwei-srv",
		Name: srvName,
		Port: portNum,
		//Address: "192.168.1.156",
		Check: &api.AgentServiceCheck{
			CheckID:                        "111",
			Name:                           "srv-chk-name",
			Interval:                       "2s",
			//TTL:                            "15m",
			//Timeout:                        "",
			HTTP:                           "http://192.168.1.156:" + port,
			DeregisterCriticalServiceAfter: "10s",
		},
	}
	logger.Infof("register to consul")
	//注册到consul
	if err := a.sd.Agent().ServiceRegister(reg); err != nil {
		log.Fatal("register failure: ", err)
	}
}

func (a *application) srvStart() {
	logger.Infof("new http server with echo")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	a.srv = e
}

func main() {
	app := newDefaultApp( "jw-echo-9000")
	app.SrvRigister("jw.health-1", "9000")
	app.srvStart()

	//app.sd.Health().Service()
	logger.Errorf("%v", app.srv.Start(":9000"))
	//go logger.Errorf("%v", app.srv.Start(":9000"))

	//app2 := newDefaultApp( "jw-echo-9001")
	//app2.SrvRigister("jw.health-2", "9001")
	//app2.srvStart()
	//
	////app.sd.Health().Service()
	//go logger.Errorf("%v", app.srv.Start(":9001"))
}

