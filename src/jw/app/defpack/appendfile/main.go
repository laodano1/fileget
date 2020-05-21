package main

import (
	"fileget/util"
	"fmt"
	"os"
	"time"
)

func main() {
	dir, _ := util.GetFullPathDir()
	f,  _ := os.OpenFile(fmt.Sprintf("%v%ctest-%v.txt", dir, os.PathSeparator, 1), os.O_APPEND | os.O_CREATE | os.O_RDWR, 0744)
	tk := time.Tick(2 * time.Second)
	for _ = range tk {
		n, err := f.WriteString("00000000000000000000000000000")
		if err != nil {
			util.Lg.Errorf("write string error: %v", err)
		}

		util.Lg.Infof("write num: %v", n)

	}


}
