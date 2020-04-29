package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"os"
	"path/filepath"
	"strings"
)

var (
	lg = golog.New("my-backend")
	exeAbsPath string
)

type oneType struct {
	typeName string
	names  []string
	paths   []string
}


type filesList struct {
	list map[string]*oneType
}

func mywalkfunc(fl filesList) func(path string, info os.FileInfo, err error) error {
	setVal := func(typeStr, path string) {
		tpss := strings.Split(path, fmt.Sprintf("%c", os.PathSeparator))
		if _, ok := fl.list[typeStr]; !ok {  // 不存在该type 才新建
			ot := &oneType{typeName: typeStr}
			ot.names = append(ot.names, tpss[len(tpss) - 1])
			ot.paths = append(ot.paths, fmt.Sprintf("/%v", tpss[len(tpss) - 1]))
			fl.list[typeStr] = ot
		} else {
			fl.list[typeStr].names = append(fl.list[typeStr].names, tpss[len(tpss) - 1])
			fl.list[typeStr].paths = append(fl.list[typeStr].paths, fmt.Sprintf("/%v", tpss[len(tpss) - 1]))
		}
	}

	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			tmpStrs := strings.Split(path, ".")
			switch tmpStrs[len(tmpStrs) - 1] {
			case "mp3":
				setVal("mp3", path)
			case "mp4":
				setVal("mp4", path)
			case "mkv":
				setVal("mkv", path)
			case "rm":
				setVal("rm", path)
			case "rmvb":
				setVal("rmvb", path)
			case "mov":
				setVal("mov", path)
			case "wmv":
				setVal("wmv", path)
			case "flv":
				setVal("flv", path)
			case "avi":
				setVal("avi", path)
			case "3gp":
				setVal("3gp", path)
			default:
			}
		}
		return err
	}
}

func main() {
	dir := "E:\\迅雷下载"

	fl := filesList{}
	fl.list = make(map[string]*oneType)

	filepath.Walk(dir, mywalkfunc(fl))

	for tp, list := range fl.list {
		lg.Debugf("type: %v", tp)
		for idx, name := range list.names {
			lg.Debugf("       file name: %v", name)
			lg.Debugf("       file path: %v", list.paths[idx])
		}
	}
}
