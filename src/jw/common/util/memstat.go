package util

import (
	"github.com/davyxu/golog"
	"runtime"
	"time"
)


var (
	//logger *golog.Logger

)

func bToM(b uint64) uint64 {
	return b / 1024
}

func ShowMemStat(d time.Duration, logger *golog.Logger) {
	tk := time.Tick(d * time.Second)
	memStat := &runtime.MemStats{}
	runtime.ReadMemStats(memStat)
	for {
		select {
		case <-tk:
			logger.Infof("Alloc: %v Kb, Totoal Alloc: %v Kb", bToM(memStat.Alloc), bToM(memStat.TotalAlloc))
		}
	}
}
