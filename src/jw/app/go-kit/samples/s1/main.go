package main

import "fileget/util"

func main() {
	mp := make(map[string]bool)
	mp["1"] = true
	mp["2"] = true
	mp["3"] = true
	util.Lg.Debugf("len: %v", len(mp))

	delete(mp, "2")
	util.Lg.Debugf("len: %v", len(mp))

}
