package main

import (
	"bytes"
	"fileget/util"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	lth = 10 * 1024 * 1024

)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("main go routine panic: %v", e)
		}
	}()
	num := 1000
	dir, _ := util.GetFullPathDir()
	wg := &sync.WaitGroup{}
	b := make([]byte, lth)
	for i := 0; i < lth; i++ {
		b = append(b, 0)
	}
	for i := 0; i < num; i++ {
		go func(id int) {
			wg.Add(1)
			// open output file
			fo, err := os.OpenFile(fmt.Sprintf("%v%ctst%ctest-%v.txt", dir, os.PathSeparator, os.PathSeparator, id), os.O_APPEND | os.O_RDWR | os.O_CREATE, 0777)
			//fo, err := os.Create(fmt.Sprintf("%v%ctst%ctest-%v.txt", dir, os.PathSeparator, os.PathSeparator, id))
			if err != nil {
				panic(err)
			}
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("goroutine panic: %v", e)
				}
			}()
			// close fo on exit and check for its returned error
			defer func() {
				if err := fo.Close(); err != nil {
					panic(err)
				}
			}()

			bf := bytes.NewBuffer(b)
			for _ = range time.Tick(2 * time.Second) {
				//n, err := bf.WriteTo(fo)
				n, err := fo.WriteString(bf.String())
				if err != nil {
					util.Lg.Errorf("(%v): write error :%v", id, err)
				} else {
					util.Lg.Infof("(%v): write byte num: %v", id, n)
				}
			}
		}(i)
	}

	wg.Wait()

	util.Lg.Debugf("bye bye!")

}
