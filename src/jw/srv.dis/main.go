package main

import (
	"github.com/davyxu/golog"
	"github.com/hashicorp/consul/api"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type (

	application struct {
		Name string
		sd *api.Client
		srv *echo.Echo
	}

)

var logger = golog.New("service.discovery")

func newDefaultApp() (app *application) {
	app = new(application)
	config := api.DefaultConfig()
	config.Address = "192.168.1.178:8500"

	logger.Infof("New consul client")
	client, err := api.NewClient(config)
	if err != nil {
		logger.Errorf("new client error: %v", err)
	}

	app.Name = "jw-echo-9000"
	app.sd = client
	logger.Infof("New client agent")
	//agent := client.Agent()


	return
}

func (a *application) SrvRigister() {
	//注册到consul的服务器内容
	logger.Infof("fill in register contents")
	reg := &api.AgentServiceRegistration{
		ID: "jinwei-srv-:5000",
		Name: "jw.health.v1",
		Port: 9000,
		Address: "192.168.1.156",
		Check: &api.AgentServiceCheck{
			CheckID:                        "111",
			Name:                           "srv-chk-name",
			Interval:                       "2s",
			//TTL:                            "15m",
			//Timeout:                        "",
			HTTP:                           "http://192.168.1.156:9000",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	logger.Infof("register to consul")
	//注册到consul
	if err := a.sd.Agent().ServiceRegister(reg); err != nil {
		log.Fatal("register failure:", err)
	}
	//time.Sleep(1 * time.Minute)
	//logger.Infof("Bye bye!")
}

func (a *application) srvStart() {
	logger.Infof("new http server with echo")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	a.srv = e
	//return
}

func main() {
	app := newDefaultApp()
	app.SrvRigister()
	app.srvStart()

	//app.sd.Health().Service()

	logger.Errorf("%v", app.srv.Start(":9000"))
}

