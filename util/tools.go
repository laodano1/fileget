package util

import (
"encoding/json"
	"errors"
	"fmt"
"github.com/davyxu/golog"
"io/ioutil"
"os"
"path/filepath"
	"strconv"
	"strings"
)

var (
	Lg = golog.New("utils")

)

func init() {
	Lg.EnableColor(true)
	Lg.SetParts(golog.LogPart_Level, golog.LogPart_TimeMS, golog.LogPart_ShortFileName)
}

func GetFullPath() (exeAbsPath string, err error) {
	exeAbsPath, err = os.Executable()
	if err != nil {
		Lg.Errorf("get os.Executable failed: %v", err)
		return
	}
	//
	Lg.Infof("os.Executable path: %v\n", exeAbsPath)

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
		Lg.Errorf("json MarshalIndent failed: %v", err)
		return
	}

	err = ioutil.WriteFile(fmt.Sprintf("%v", fileName), f, 0644)
	if err != nil {
		Lg.Errorf("write to json file failed: %v", err)
		return
	}

}

func GetMatchedFiles(dir, suffix string) (matchedFiles map[string]bool) {
	matchedFiles = make(map[string]bool)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) (er error) {
		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			matchedFiles[info.Name()] = true
		}
		return
	})

	return
}


func GetFiles(dir, suffix string) (files []string, er error) {
	er = filepath.Walk(dir, func(path string, info os.FileInfo, err error) (er error) {
		defer func() {
			if err := recover(); err != nil {

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

func ReadWHLJson(fileName string) (jsonb []byte, err error) {
	jsonb, err = ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	return
}


func ValueOfString(input string) int {
	if len(input) == 0 {return 0}

	rs := []rune(input)
	sb := make([]rune, 0)
	nbs := make([][]string, 0)
	for i, j := 0, 0; i < len(input); i++ {
		switch rs[i] {
		case ' ':
			continue
		case '+', '-', '*':
			sb = append(sb, rs[i])
		default:
			if i == 0 {
				tmp := make([]string, 0)
				tmp = append(tmp, string(rs[i]))
				nbs = append(nbs, tmp)
			} else {
				if rs[i-1] == '+' || rs[i-1] == '-' ||rs[i-1] == '*' || rs[i-1] == ' ' {
					tmp := make([]string, 0)
					tmp = append(tmp, string(rs[i]))
					nbs = append(nbs, tmp)
					j++
				}  else {
					nbs[j] = append(nbs[j], string(rs[i]))
				}
			}
		}
	}
	//util.Lg.Debugf("sb: %c", sb)
	//util.Lg.Debugf("nbs: %v", nbs)

	var sum int
	nums := make([]int, 0)
	for i := range nbs {
		t1, _ := strconv.Atoi(strings.Join(nbs[i], ""))
		nums = append(nums, t1)
	}

	for i, j := 0, 0; i < len(nums); i++ {
		if i == 0 {
			sum = nums[i]; continue
		} else {
			sum = calculate(sum, nums[i], sb[j])
			//util.Lg.Debugf("sum: %v", sum)
			j++
		}
	}

	return sum
}

func calculate(a, b int, s rune) (sum int) {
	if s == '+' {
		sum = a + b
	} else if s == '-' {
		sum = a - b
	} else {
		sum = a * b
	}
	return
}

