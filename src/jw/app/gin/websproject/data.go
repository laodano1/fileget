package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirList       dirsJsonObj

)


func mywalkfunc(fl filesList) func(path string, info os.FileInfo, err error) error {
	setVal := func(typeStr, path string) {
		tpss := strings.Split(path, fmt.Sprintf("%c", os.PathSeparator))
		if _, ok := fl.list[typeStr]; !ok {  // 不存在该type 才新建
			ot := &oneType{typeName: typeStr}
			ot.names = append(ot.names, tpss[len(tpss) - 1])
			ot.paths = append(ot.paths, fmt.Sprintf("/video/%v", tpss[len(tpss) - 1]))
			fl.list[typeStr] = ot
		} else {
			fl.list[typeStr].names = append(fl.list[typeStr].names, tpss[len(tpss) - 1])
			fl.list[typeStr].paths = append(fl.list[typeStr].paths, fmt.Sprintf("/video/%v", tpss[len(tpss) - 1]))
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

func LoadMediaInfo() (err error) {
	c, err := ioutil.ReadFile("dirs.json")
	if err != nil {
		lg.Errorf("read config file failed: %v", err)
		return
	}

	err = json.Unmarshal(c, &dirList)
	if err != nil {
		lg.Errorf("Unmarshal to dir list failed: %v", err)
		return
	}

	var videoFileList filesList
	videoFileList.list = make(map[string]*oneType)
	for _, dirItem := range dirList.DirList {
		filepath.Walk(dirItem, mywalkfunc(videoFileList))
	}

	setWebData(videoFileList)
	return
}

func setWebData(videos filesList) {
	sbObj := &sidebar{}
	sbmi := sidebarMainItem{}

	subIList := make([]sbSubItem, 0)
	for tp, _ := range videos.list {
		// sidebar sub-item
		subIList = append(subIList, sbSubItem{Name: tp, Href: fmt.Sprintf("/%v", tp)})
	}
	sbmi.SubItems = subIList

	pgcList := make([]pageContent, 0)
	for _, atype := range videos.list {
		// sub-item relevant page items
		pgItems := make([]pageItem, 0)
		for idx, name := range atype.names {
			pgItems = append(pgItems, pageItem{Name: name, Href: atype.paths[idx]})
		}

		pgcList = append(pgcList, pageContent{PageObjs: pgItems})
	}
	sbmi.PageCnt  = pgcList
	sbmi.Name = "Media"
	sbObj.List = append(sbObj.List, sbmi)

	sidebarData = *sbObj
}
