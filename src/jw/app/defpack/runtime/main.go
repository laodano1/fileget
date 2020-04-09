package main

import (
	"crypto/md5"
	"github.com/davyxu/golog"
	"time"
	"unsafe"
)

func main() {
	logger := golog.New("haha")

	arr := []int64{1, 2, 3, 4}
	logger.Infof("arr occupies: %v bytes", unsafe.Sizeof(arr))

	b := []byte("aabbcc")
	s :=time.Now()
	cs := md5.Sum(b)
	t := time.Now().Sub(s).Nanoseconds()
	logger.Infof("take time: %v(nanoseconds), to get md5 sum: %v", t, cs)

	
}
