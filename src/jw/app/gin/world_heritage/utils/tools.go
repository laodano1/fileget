package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetFullPath() (exeAbsPath string, err error) {
	exeAbsPath, err = os.Executable()
	if err != nil {
		fmt.Errorf("get os.Executable failed: %v", err)
		return
	}
	//
	fmt.Printf("os.Executable path: %v", exeAbsPath)

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

func ReadJson(exeAbsPath, fileName string) (jsonb []byte, err error) {
	jsonb, err = ioutil.ReadFile(exeAbsPath + "/public/" + fileName)
	if err != nil {
		return
	}
	return
}

func ReadWHLJson(fileName string) (jsonb []byte, err error) {
	jsonb, err = ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	return
}

func GetFiles(dir, suffix string) (files []string, er error) {
	er = filepath.Walk(dir, func(path string, info os.FileInfo, err error) (er error) {
		defer func() {
			if err := recover(); err != nil {
				//fmt.Println(err)
				er = errors.New(fmt.Sprintf("%v", err))
				return
			}
		}()
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			files = append(files, path)
		}

		return err
	})
	return
}



//
//
//func UnmarshalJson(jsonb []byte, val interface{}) (err error) {
//	err = json.Unmarshal(jsonb, val)
//	if err != nil {
//		return
//	}
//	return
//}