package main

import (
	"fileget/util"
	"time"
)

func main() {

	util.ZLogger.Infof("hell world")
	util.ZLogger.Errorf("hell world")
	util.ZLogger.Debugf("hell world")

	util.ZLogger.Debugf(time.Unix(1607350510, 0).Format(time.RFC3339Nano))
	util.ZLogger.Debugf(time.Unix(1607351868, 0).Format(time.RFC3339Nano))

}
