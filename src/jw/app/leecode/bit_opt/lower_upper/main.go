package main

import "fileget/util"

func main() {
	t := 'Z'

	tt := t | ' '  // 与英文空字符或操作 转换为小写

	t1 := 'c' & '_'  // 与下划线(underscore)与操作 转换为大写

	t2 := t ^ ' '  // 与英文空字符异或 大小写互换

	util.Lg.Debugf("t1: %c, tt: %v, t2: '%c'", t1, tt, t2)

	t2 = t2 ^ ' '  // 与空字符异或 大小写互换
	util.Lg.Debugf("tt: %c, t2: '%c'", tt, t2)
}
