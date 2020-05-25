package main

import (
	"fileget/util"
	"net"
	"os"
)

func getGatewayIP() {

	host, _ := os.Hostname()
	util.Lg.Infof("1. host: %v", host)
	//
	cn, _ := net.LookupMX(host)
	util.Lg.Infof("1. LookupCNAME: %v", cn)
	//addrs, _ := net.LookupIP(host)
	//for _, addr := range addrs {
	//	addr.
	//}

}

func getLocalHostIP() string {
	//conn, err := net.Dial("udp", "8.8.8.8:80")
	conn, err := net.Dial("udp", "10.0.0.1:53")
	if err != nil {
		util.Lg.Errorf("dial faile: %v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	util.Lg.Debugf("LocalAddr: %v", localAddr.IP)
	//ip := localAddr.IP.String()
	return localAddr.IP.String()
}

func main() {
	getLocalHostIP()
	getGatewayIP()
}
