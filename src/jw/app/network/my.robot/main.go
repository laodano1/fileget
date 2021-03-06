package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"runtime"
	"sort"

	_ "github.com/davyxu/cellnet/peer/gorillaws"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

var (
	srvAdd = ":8888"
	logger *golog.Logger
	memStat *runtime.MemStats
)


//
func FeijiPokerParse(pokers []int32, daiPCnt int) (zhuP [][]int32, daiP [][]int32) {
	pokerMap := make(map[int][]int32)
	for idx, v := range pokers {
		pm := v & 0x0F   //牌面值
		pokerMap[int(pm)] = append(pokerMap[int(pm)], pokers[idx])
	}

	keys := make([]int, 0)
	//把单牌取出后，删除该节点
	for k, v := range pokerMap {
		//带牌个数，带单，带双
		if daiPCnt != 0 {
			if len(v) == daiPCnt {
				daiP = append(daiP, v)
				delete(pokerMap, k)
			} else {
				keys = append(keys, k)
			}
		}
	}

	sort.Ints(keys)
	zhuP = make([][]int32, 0)
	//fmt.Printf("value: %v , len zhuP: %v", zhuP, len(zhuP))
	//zhuP = append(zhuP[1:], )
	for _, v := range keys {
		zhuP = append(zhuP, pokerMap[v])
	}
	if daiPCnt == 0 {
		daiP = make([][]int32, 0)
	}

	return
}

var (
	huaSe = []string{
		"方块",
		"梅花",
		"红桃",
		"黑桃",
		"王",
	}

	mianZhi = []string{
		"0",
		"A",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"J",
		"Q",
		"K",
	}
)


//生成玩家视角的牌面
func GetPaiValue(pais []int32) []string {
	paiView := make([]string, 0)
	for _, v := range pais {
		hs := 0xF0 & v
		hs = hs >> 4    //花色
		val := 0x0F & v //面值

		if hs == 4 { //王牌
			switch val {
			case 0x0E:
				paiView = append(paiView, "小王")
			case 0x0F:
				paiView = append(paiView, "大王")
			}
		} else {
			paiView = append(paiView, fmt.Sprintf("%v%v", huaSe[hs], mianZhi[val]))
		}

	}
	return paiView
}

func main()  {
	logger = golog.New("my.robot")
	logger.SetParts()

	//curp := []int32{0x03, 0x13, 0x23, 0x33, 0x08, 0x18, 0x28, 0x38}
	curp := []int32{21, 29, 38}

	fmt.Println(curp)
	fmt.Println(GetPaiValue(curp))

	//siPsR, daiP := 	FeijiPokerParse(curp, 2)
	//
	//fmt.Println("主牌:", siPsR)
	//fmt.Println(daiP)



	//go util.ShowMemStat(10, logger)
	//
	//queue := cellnet.NewEventQueue()
	//
	//p := peer.NewGenericPeer("gorillaws.Connector", "my-robot-cli", srvAdd, queue)
	//
	//proc.BindProcessorHandler(p, "gorillaws.ltv", func(ev cellnet.Event) {
	//	switch msg := ev.Message().(type) {
	//	case *cellnet.SessionConnected:
	//		logger.Infof("session(%v) connected, %v", ev.Session().ID(), msg.SystemMessage)
	//
	//	case *cellnet.SessionClosed:
	//		logger.Infof("Session(%v) closed", ev.Session().ID())
	//	}
	//})
	//
	//p.Start()
	//
	//queue.StartLoop()
	//
	//queue.Wait()

}



