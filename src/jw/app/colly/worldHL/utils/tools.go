package utils

import (
	"encoding/json"
	"github.com/davyxu/golog"
	"io/ioutil"
)

var (
	lg = golog.New("utils")
)

func Write2JsonFile(whl interface{}, fileName string) {

	f, err := json.MarshalIndent(whl, "", "  ")
	if err != nil {
		lg.Errorf("json MarshalIndent failed: %v", err)
		return
	}

	err = ioutil.WriteFile(fileName, f, 0644)
	if err != nil {
		lg.Errorf("write to json file failed: %v", err)
		return
	}

}
