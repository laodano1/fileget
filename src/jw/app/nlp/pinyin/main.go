package main

import (
	"fileget/src/jw/app/nlp/pinyin/lib"
	"flag"
)

func main() {
	//util.Lg.Debugf("pin yin: %v", pinyin.Pinyin(hans, a))
	flag.StringVar(&lib.Addr, "p", lib.Addr, "listen port setting")
	flag.Parse()

	lib.StartServer()

}
