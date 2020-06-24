package main

import (
	"fileget/util"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	//hans := "范蠡"
	hans := "弍是"
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone

	util.Lg.Debugf("pin yin: %v", pinyin.Pinyin(hans, a))

}
