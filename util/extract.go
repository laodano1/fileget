package util

import (
	"fmt"
	"strings"
	"os"
	"time"
	"sync"
	"os/exec"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"regexp"
)

type TypeItem struct {
	Type     string `yaml: "type"`
	Name     string `yaml: "name"`
	Item     string `yaml: "item"`
}

type Type map[string][]TypeItem
	//
//type Collections struct {
//	Interval []string `yaml: "interval"`
//	Duration []string `yaml: "duration"`
//	//Types    map[string][]Type `yaml: "Collections"`
//	Types    []interface{} `yaml: "Collections"`
//}
type Collections map[string][]TypeItem
//type Collections map[string][]interface{}

	// store local config file
type Resources struct {
	Interval time.Duration
	Nodes    []string
	Pods     []string
	Wg       sync.WaitGroup
	Duration time.Duration
	Debug    bool
	CFileLoc string
	OpFileloc string
}

const (
	cmd    = "kubectl"
)

var (
	//col Collections
	col = Collections{}
	//cfileloc = "C:\\Users\\jinwu\\go\\src\\jwCmdApp\\cm.yaml"
	cfileloc = "./cm.yaml"
)


// check file in path exists or not
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsExist(err)
}

func (r *Resources) TearDown()  {
	fmt.Println("wait monitor workers start")
	r.Wg.Wait()
	fmt.Println("wait monitor workers end\n\nBye bye!")
}

func (r *Resources) Init() error {

	yamlFile, err := ioutil.ReadFile(cfileloc)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v\n", err)
		return err
	}

	//fmt.Println(string(yamlFile))

	//err = yaml.Unmarshal(yamlFile, &tst)
	err = yaml.Unmarshal(yamlFile, &col)
	if err != nil {
		fmt.Printf("yaml unmarshal failed! %s\n", err)
		return err
	} else {
		fmt.Println("yaml unmarshal successfully!")
	}

	//fmt.Println(col)

	//get interval
	itv, err := strconv.Atoi(col["interval"][0].Item)
	if err != nil {
		fmt.Println("convert interval error:", err)
	}
	r.Interval = time.Duration(itv) * time.Second

	// get duration
	drt, err := strconv.Atoi(col["duration"][0].Item)
	if err != nil {
		fmt.Println("convert interval error:", err)
	}
	r.Duration = time.Duration(drt * 30)  * time.Second

	r.CFileLoc  = col["confileloc"][0].Item
	r.OpFileloc = col["outputloc"][0].Item

	// get debug
	if col["debug"][0].Item == "1" {
		r.Debug = true
	} else {
		r.Debug = false
	}


	fmt.Printf("interval: %v\n", r.Interval)
	fmt.Printf("duration: %v\n", r.Duration)
	fmt.Printf("cfile location: %v\n", r.CFileLoc)
	fmt.Printf("output location: %v\n", r.OpFileloc)

	if FileExists(r.CFileLoc) {
		fmt.Println(r.CFileLoc, "do exist!")

		// extract data from it
		r.ExtractConfig(r.CFileLoc)
	} else {
		fmt.Println(r.CFileLoc, "doesn't exist!")
	}

	return nil
	//fmt.Printf("%s\n%s\n", res.Nodes, res.Pods)
}

func ReplMB(src *string, regex string, repl string) error {
	r, err := regexp.Compile(regex)
	if err != nil {
		fmt.Printf("complie regexp error %s\n", err)
		return err
	}
	*src = r.ReplaceAllString(*src, repl)
	//fmt.Println(src)
	return nil
}

func IsEmptyFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		if !strings.ContainsAny(err.Error(), "no such file") {
			fmt.Printf("stat file %s error: %s\n", path, err)
			return false
		} else {
			return true
		}
	}
	if f.Size() != 0 {
		return false
	}
	return true
}

func (r *Resources) CmdOutput(list []string, cmd string)  {
	for i, v := range list {
		//r.Wg.Add(1)
		go func(id int, name string) {
			var outpfExist bool
			var outputFile string = r.OpFileloc + name + ".csv"

			t := time.NewTicker(r.Interval)

			if IsEmptyFile(outputFile) {
				outpfExist = false
			} else {
				outpfExist = true
			}

			f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Create file %s failed!\n", f.Name())
			}
			tm := time.NewTimer(r.Duration)

			if !outpfExist {   // only not exist, then add title
				if strings.HasPrefix(name, "isdk-cluster") {  //node
					f.WriteString("TIME,NAME,CPU(cores),CPU%,MEMORY(bytes),MEMORY%\n")
				} else {  // pod
					f.WriteString("TIME,NAME,CPU(cores),MEMORY(bytes)\n")
				}
			}

			for {
				select {
				case <-t.C:
					var kubeOut []byte
					//var err error
					if r.Debug {
						kubeOut = []byte(time.Now().Format(time.RFC3339Nano))
					} else {
						var kubeCmd *exec.Cmd
						if strings.HasPrefix(name, "isdk-cluster") {
							kubeCmd = exec.Command("kubectl", "top", "node", name)
						} else {
							kubeCmd = exec.Command("kubectl", "top", "pod", name)
							//fmt.Printf("kubectl top pod %s\n", name)
						}

						kubeOut, err = kubeCmd.Output()
						if err != nil {
							fmt.Printf("cmd failed: %s!\n", err)
						}
					}

					regexp.Compile("[ ]+")
					str := strings.Split(string(kubeOut), "\n")[1]
					ReplMB(&str, "[ ]+", " ")
					ReplMB(&str, " ", ",")
					str = str[:len(str) -1]

					now := time.Now().Format(time.RFC3339Nano)
					f.WriteString(now + "," + string(str) + "\n")

					fmt.Printf("%55v | #%d worker gets top info\n", time.Now(), id)

				case <- tm.C:
					fmt.Printf("time is up, #%d monitor worker ends!\n", id )
					f.Close()
					r.Wg.Done()
					return
				}
			}
		}(i, v)
	}
}

func (r *Resources) MonitorKubeOutput(rType string)  {
	fmt.Printf("in MonitorKubeOutput type: %s\n", rType)
	switch strings.ToLower(rType) {
	case "all" :
		r.Nodes = append(r.Nodes, r.Pods...)
		r.Wg.Add(len(r.Nodes))
		r.CmdOutput(r.Nodes, cmd)
	case "node" :
		r.Wg.Add(len(r.Nodes))
		r.CmdOutput(r.Nodes, cmd)
	case "pod" :
		r.Wg.Add(len(r.Pods))
		r.CmdOutput(r.Pods, cmd)
	default:

	}
}

// extract data from cm.yaml
func (r *Resources) ExtractConfig(path string) error {
	//
	for _, tpItem := range col["collections"] {
		var tmp TypeItem

		tmp.Name = tpItem.Name
		tmp.Item = tpItem.Item
		//fmt.Println(tpItem)
		//fmt.Printf("%s:%s", tmp.TypeName, tmp.Item)

		switch strings.ToUpper(tmp.Name) {
		case "POD" :
			pods := strings.Split(tmp.Item, " ")
			for _, it := range pods {
				//fmt.Printf("pod: '%s' will be monitored!\n", it)
				r.Pods = append(r.Pods, it)
			}

		case "NODE" :
			if strings.ToUpper(tmp.Item) == "ALL" {
				//fmt.Println("'All' Node will be monitored!")
				//nodeChkFlag = true
			} else {
				//fmt.Printf("node: '%s' will be monitored!", tmp.Item)
				nodes := strings.Split(tmp.Item, " ")
				for _, it := range nodes {
					//fmt.Printf("node '%s' will be monitored!\n", it)
					r.Nodes = append(r.Nodes, it)
				}
			}

		default:
			fmt.Println("bad type value in cm.yaml! Only node/pod are supported!")
		}
	}

	//fmt.Printf("%s\n", r.Nodes)
	//fmt.Printf("%s\n", r.Pods)
	//fmt.Println(r)

	return nil
}
