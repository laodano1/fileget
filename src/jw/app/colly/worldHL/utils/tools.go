package utils

import (
	"encoding/json"
	"fmt"
	"github.com/davyxu/golog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	lg = golog.New("utils")
)

func GetFullPath() (exeAbsPath string, err error) {
	exeAbsPath, err = os.Executable()
	if err != nil {
		fmt.Errorf("get os.Executable failed: %v", err)
		return
	}
	//
	fmt.Printf("os.Executable path: %v\n", exeAbsPath)

	return
}

func GetFullPathDir() (exeAbsPathDir string, err error) {
	exeAbsPathDir, err = GetFullPath()
	if err != nil {
		return
	}
	exeAbsPathDir = filepath.Dir(exeAbsPathDir)

	return
}

func Write2JsonFile(whl interface{}, fileName string) {

	f, err := json.MarshalIndent(whl, "", "  ")
	if err != nil {
		lg.Errorf("json MarshalIndent failed: %v", err)
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%v", fileName), f, 0644)
	if err != nil {
		lg.Errorf("write to json file failed: %v", err)
		return
	}

}

func GetMatchedFiles(dir, suffix string) (matchedFiles map[string]bool) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) (er error) {
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			matchedFiles[info.Name()] = true
		}
		return
	})

	return
}
