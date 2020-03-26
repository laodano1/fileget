package main

import (
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"google.golang.org/grpc"
	"net/http"
)

func aaa(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}

type myService struct {
	httpSvc web.Service
	grpcSvc grpc.Server
}

func main() {
	lg := golog.New("ttttt")
	ms := myService{}

	ms.httpSvc = web.NewService(
		web.Address(":3000"),
		web.Registry(consul.NewRegistry(registry.Addrs("10.0.0.131:8501"))),
	)
	router := gin.New()
	router.GET("/", aaa)
	ms.httpSvc.Handle("", router)

	lg.Debugf("have started http svc")

	ms.grpcSvc = grpc.NewServer()
	ms.grpcSvc.
		ms.grpcSvc.Start()

	lg.Debugf("have started http svc 2")

	ms.httpSvc.Run()
	//svc := micro.NewService(
	//	micro.Name("test-2"),
	//	micro.Address(":8888"),
	//	micro.Registry(
	//		consul.NewRegistry(
	//			registry.Addrs("10.0.0.131:8501"),
	//		),
	//	),
	//	micro.RegisterTTL(2*time.Second),
	//)
	//
	//if err := svc.Run(); err != nil {
	//	lg.Errorf("micro service run failed: %v", err)
	//}
}
