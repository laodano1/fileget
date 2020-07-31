package main

import "fileget/util"

func main() {
	t := 'Z'

	tt := t | ' '

	util.Lg.Debugf("a: '%v', %b", 'a', 'a')
	util.Lg.Debugf("A: '%v', %b", 'A', 'A')
	util.Lg.Debugf("z: '%v', %b", 'z', 'z')
	util.Lg.Debugf("Z: '%v', %b", 'Z', 'Z')
	util.Lg.Debugf("' ' : '%v', %b", ' ', ' ')
	util.Lg.Debugf("tt: '%c'", tt)

}
