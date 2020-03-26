package main

import (
	"github.com/davyxu/golog"
	"time"
)

func main() {
	lg := golog.New("mylist")

	//r := ring.New(1)
	//lg.Debugf("1 list len: %v", r.Len())
	//for _, v := range []int{1, 2, 3, 4, 5} {
	//	item := &ring.Ring{
	//		Value: func() int {
	//			return v
	//		},
	//	}
	//	r.Link(item)
	//}
	//
	//lg.Debugf("2 list len: %v", r.Len())
	lg.Debugf("now: %v", time.Now().Format(time.RFC3339Nano))
}
