module fileget

go 1.13

replace (
	github.com/kubernetes-incubator/service-catalog => github.com/kubernetes-sigs/service-catalog v0.3.0
	github.com/kubernetes/kubernetes => k8s.io/kubernetes v1.18.5
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200214081623-ecbd4af0fc33
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200214081019-7490b3ed6e92
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200214082307-e38a84523341
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200214080538-dc8f3adce97c
)

// k8s relevant
require (
	//github.com/kubernetes/kubernetes v1.18.5
	github.com/containernetworking/cni v0.8.0 // indirect
	github.com/kubernetes-incubator/service-catalog v0.3.0 // indirect
	github.com/projectcalico/cni-plugin v3.8.9+incompatible // indirect
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/kube-aggregator v0.18.5 // indirect
	k8s.io/kube-controller-manager v0.18.5 // indirect
	k8s.io/kube-proxy v0.18.5
	k8s.io/kube-scheduler v0.18.5 // indirect
	k8s.io/kubectl v0.18.5
	k8s.io/kubelet v0.18.5 // indirect
	k8s.io/legacy-cloud-providers v0.18.5
	k8s.io/utils v0.0.0-20200410165547-614e4363e9c4 // indirect
)

require (
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
)

require (
	github.com/docker/docker-ce v17.12.1-ce-rc2+incompatible // indirect
	github.com/gin-gonic/gin v1.6.2
)

require (
	github.com/99designs/gqlgen v0.9.1
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/benmoss/go-powershell v0.0.0-20190925205200-09527df358ca // indirect
	github.com/blackjack/webcam v0.0.0-20200313125108-10ed912a8539
	github.com/bronze1man/goStrongswanVici v0.0.0-20200615065859-ff78c6a7cf1f // indirect
	github.com/chromedp/cdproto v0.0.0-20200424080200-0de008e41fa0
	github.com/chromedp/chromedp v0.5.3
	github.com/chromedp/examples v0.0.0-20200501161515-cb21abae103c // indirect

	github.com/coreos/flannel v0.12.0 // indirect
	github.com/coreos/go-iptables v0.4.5 // indirect
	github.com/davyxu/cellnet v4.1.0+incompatible
	github.com/davyxu/golog v0.1.0
	github.com/davyxu/goobjfmt v0.1.0 // indirect
	github.com/davyxu/protoplus v0.1.0 // indirect
	github.com/denverdino/aliyungo v0.0.0-20200701124158-2451fe6f9270 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible

	github.com/gizak/termui v3.1.0+incompatible // indirect
	github.com/gizak/termui/v3 v3.1.0 // indirect
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gocolly/colly/v2 v2.0.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/gwuhaolin/livego v0.0.0-20200509033525-ac935b8214d0 // indirect
	github.com/hailocab/go-geoindex v0.0.0-20160127134810-64631bfe9711
	github.com/hashicorp/consul/api v1.3.0
	github.com/hashicorp/go-uuid v1.0.1
	github.com/headzoo/surf v1.0.0 // indirect
	github.com/headzoo/ut v0.0.0-20181013193318-a13b5a7a02ca // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/juju/errors v0.0.0-20200330140219-3fe23663418f // indirect
	github.com/juju/testing v0.0.0-20200706033705-4c23f9c453cd // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/magicsea/behavior3go v0.0.0-20200226033918-88f465325648
	github.com/micro/cli v0.2.0
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.1.2-0.20190710094942-bf407858372c
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro v1.7.1-0.20190711215914-2cddc2c877c5
	github.com/micro/micro/v2 v2.7.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mozillazg/go-pinyin v0.18.0
	github.com/nats-io/nats.go v1.9.2
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1 // indirect
	github.com/pborman/uuid v1.2.0

	github.com/rakelkar/gonetsh v0.0.0-20190930180311-e5c5ffe4bdf0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/tealeg/xlsx v1.0.5
	github.com/valyala/fasthttp v1.12.0
	github.com/vektah/gqlparser v1.1.2
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20200520041808-52d707b772fe // indirect
	github.com/yanyiwu/gojieba v1.1.2 // indirect
	go.mongodb.org/mongo-driver v1.3.4

	gopkg.in/headzoo/surf.v1 v1.0.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8

)
