package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirList   dirsJsonObj
	staticDir []string
)

func mywalkfunc(fl filesList, staticDir, rootDir string) func(path string, info os.FileInfo, err error) error {
	setVal := func(typeStr, path string) {
		tpss := strings.Split(path, fmt.Sprintf("%c", os.PathSeparator))
		restStr := strings.Replace(path, rootDir, "", -1)
		restStr = strings.Replace(restStr, "\\", "/", -1)
		if _, ok := fl.list[typeStr]; !ok { // 不存在该type 才新建
			ot := &oneType{typeName: typeStr}
			ot.names = append(ot.names, tpss[len(tpss)-1])

			ot.paths = append(ot.paths, fmt.Sprintf("/%v/%v", staticDir, restStr))
			fl.list[typeStr] = ot
		} else {
			fl.list[typeStr].names = append(fl.list[typeStr].names, tpss[len(tpss)-1])
			fl.list[typeStr].paths = append(fl.list[typeStr].paths, fmt.Sprintf("/%v/%v", staticDir, restStr))
		}
	}

	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			tmpStrs := strings.Split(path, ".")
			switch tmpStrs[len(tmpStrs)-1] {
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
	epath, _ := os.Executable()
	lg.Debugf("exe path: %v", epath)
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
		uuiddir, err := uuid.GenerateUUID()
		if err != nil {
			lg.Errorf("generate uuid failed: %v", err)
			continue
		}
		staticDir = append(staticDir, uuiddir)
		filepath.Walk(dirItem, mywalkfunc(videoFileList, uuiddir, dirItem))
	}
	lg.Debugf("uuiddir： %v, mp4: %v", staticDir, len(videoFileList.list["mp4"].names))

	setWebData(videoFileList)
	return
}

func setWebData(videos filesList) {
	sbObj := &sidebar{}
	sbmi := sidebarMainItem{}

	subIList := make([]sbSubItem, 0)
	pgcList := make([]pageContent, 0)

	loadSbAndPageObj := func(atype *oneType, tp string) (pgItems []pageItem) {
		pgItems = make([]pageItem, 0)
		for idx, name := range atype.names {
			pgItems = append(pgItems, pageItem{Name: name, Href: atype.paths[idx]})
		}
		subIList = append(subIList, sbSubItem{Name: tp, Href: fmt.Sprintf("/%v", tp)})
		return
	}

	for tp, atype := range videos.list {
		var pgItems []pageItem
		switch tp {
		case "mp3":
			pgItems = loadSbAndPageObj(atype, tp)
		case "mp4":
			pgItems = loadSbAndPageObj(atype, tp)
		case "mkv":
			pgItems = loadSbAndPageObj(atype, tp)
		case "rm":
			pgItems = loadSbAndPageObj(atype, tp)
		case "rmvb":
			pgItems = loadSbAndPageObj(atype, tp)
		case "mov":
			pgItems = loadSbAndPageObj(atype, tp)
		case "wmv":
			pgItems = loadSbAndPageObj(atype, tp)
		case "flv":
			pgItems = loadSbAndPageObj(atype, tp)
		case "avi":
			pgItems = loadSbAndPageObj(atype, tp)
		case "3gp":
			pgItems = loadSbAndPageObj(atype, tp)
		default:
		}
		pgcList = append(pgcList, pageContent{PageObjs: pgItems})
	}
	sbmi.SubItems = subIList
	sbmi.PageCnt = pgcList
	sbmi.Name = "Media"

	sbObj.List = append(sbObj.List, sbmi)

	sidebarData = *sbObj
}
