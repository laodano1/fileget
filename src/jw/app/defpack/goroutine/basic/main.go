package main

import (
	"fileget/util"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

//func myRoutine(wg *sync.WaitGroup, mp *map[int]string)  {
func myRoutine(wg *sync.WaitGroup, mp *[]int)  {
	defer wg.Done()
	tk := time.Tick(2 * time.Second)
	tm := time.After(20 * time.Second)
	log.Printf("in go routine")
	for {
		select {
		case <- tk:
			log.Printf("mp => %v", mp)
		case <- tm:
			log.Printf("go routine return!")
			return
		}
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func getip() (ip string) {
	ips := make([]net.IP, 0)
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				//ip = v.IP
			}
			if ip.To4() != nil {
				if strings.HasPrefix(ip.To4().String(), "10.0.0.") {
					ips = append(ips, ip.To4())
				}
			}

			// process IP address
		}

	}

	for _, ip := range ips {
		log.Printf("ip: %v", ip)
	}

	return
}

type GameRoomConfigInfo struct {
	AgentId           int32  `json:"agentId"`           //业主id
	AntesAmount       int64  `json:"antesAmount"`       //底座单位(分)
	GameId            string `json:"gameId"`            //游戏id
	IsOpenAiRobot     bool   `json:"isOpenAiRobot"`     //是否开启ai机器人
	IsOpenMirrorRobot bool   `json:"isOpenMirrorRobot"` //是否开启镜像机器人
	Level             int32  `json:"level"`             //房间等级
	LimitAmount       int64  `json:"limitAmount"`       //单位(分)
	RoomId            string `json:"roomId"`            //房间id
	RoomName          string `json:"roomName"`          //房间名称
	ServiceFee        int64  `json:"serviceFee"`        //房间服务费
	VirtualGold		  int64  `json:"virtualGold"`		//虚拟金币(体验场)
	SpecialConfig     string `json:"specialConfig"`    //
}

func chanTst() {
	success := make(chan bool)

	tk := time.Tick(2 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		close(success)
	}()

	for {
		select {
		case <-success:
			util.Lg.Debugf("in success")
			return
		case <-tk:
			util.Lg.Debugf("in tick")
		}
	}

}

func main() {

	chanTst()

	//ip := GetOutboundIP()
	//getip()
	//str := "3086001010101"
	//it, err := strconv.Atoi(str[len(str)-5:])
	//if err != nil {
	//	log.Fatalf("convert error: %v", err)
	//}
	//log.Printf("str: %v", it)




	//ip.String()
	//log.Printf("ip: %v", ip)
	//mp := make(map[int]string)
	//mp := make([]int, 0)
	//
	//mp = append(mp, 1)
	////for i := 0; i < 4; i++ {
	////	//mp[i] = fmt.Sprintf("i=%v", i)
	////	mp[i] = i
	////}
	//
	//log.Printf("in main go routine")
	//wg := &sync.WaitGroup{}
	//wg.Add(1)
	//
	//go myRoutine(wg, &mp)
	//
	//time.Sleep(4 * time.Second)
	////mp[20] = 99
	//mp = append(mp, 99)
	////time.Sleep(1 * time.Second)
	////delete(mp, 0)
	//
	//wg.Wait()
	//log.Printf("bye bye!")
}
