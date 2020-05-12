package main

import "github.com/davyxu/golog"

func main() {
	exeDirPath, _ = utils.GetFullPathDir()
	lg.Debugf("--------- exe dir path: %v", exeDirPath)

	lg.SetParts(golog.LogPart_Level, golog.LogPart_Name, golog.LogPart_TimeMS)
	lg.EnableColor(true)
}
