package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

func ReadWHLJson(exeAbsPath, fileName string) (jsonb []byte, err error) {
	jsonb, err = ioutil.ReadFile(exeAbsPath + "/tmp/" + fileName)
	if err != nil {
		return
	}
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