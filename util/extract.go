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
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"context"
	"os/signal"
	"syscall"
)

type TypeItem struct {
	Type     string `yaml: "type"`
	Name     string `yaml: "name"`
	Item     string `yaml: "item"`
}

type Type map[string][]TypeItem
	//
type Collections map[string][]TypeItem

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
	Log      *log.Logger
	//LogChan  chan string
}

const (
	cmd    = "kubectl"
)

var (
	//col Collections
	col = Collections{}
	//cfileloc = "C:\\Users\\jinwu\\go\\src\\jwCmdApp\\cm.yaml"
	cfileloc = "./cm.yaml"
	dbDriver   = "mysql"
	user     = "grafana"
	password = "12345.Qwert"
	hostIP   = "127.0.0.1"
	port     = "3306"
	database = "forgrafana"
)


// check file in path exists or not
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsExist(err)
}

func (r *Resources) TearDown()  {
	r.Log.Println("wait monitor workers start")
	r.HandleSignal()

	r.Wg.Wait()
	os.Remove(r.OpFileloc + strconv.Itoa(os.PathSeparator) + strconv.Itoa(os.Getpid()) + ".pid")

	r.Log.Println("wait monitor workers end\tBye bye!")
}

func (r *Resources) Init() error {

	yamlFile, err := ioutil.ReadFile(cfileloc)
	if err != nil {
		r.Log.Printf("yamlFile.Get err   #%v\n", err)
		return err
	}

	logFile, err := os.OpenFile("monitor.log", os.O_APPEND | os.O_CREATE | os.O_WRONLY , 0644)
	if err != nil {
		r.Log.Println(err)
	}

	//r.Log.SetOutput(io.MultiWriter(os.Stdout, os.Stderr, logFile))

	//r.Log = log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate | log.Ltime | log.Lshortfile)
	r.Log = log.New(logFile, "", log.Ldate | log.Ltime | log.Lshortfile)

	//fmt.Println(string(yamlFile))

	//err = yaml.Unmarshal(yamlFile, &tst)
	err = yaml.Unmarshal(yamlFile, &col)
	if err != nil {
		r.Log.Printf("yaml unmarshal failed! %s\n", err)
		return err
	} else {
		r.Log.Printf("yaml unmarshal successfully!")
	}

	//fmt.Println(col)

	//get interval
	itv, err := strconv.Atoi(col["interval"][0].Item)
	if err != nil {
		r.Log.Println("convert interval error:", err)
	}
	r.Interval = time.Duration(itv) * time.Second

	// get duration
	drt, err := strconv.Atoi(col["duration"][0].Item)
	if err != nil {
		r.Log.Println("convert interval error:", err)
	}
	//r.Duration = time.Duration(drt * 30 )  * time.Second
	r.Duration = time.Duration(drt)  * time.Minute

	r.CFileLoc  = col["confileloc"][0].Item
	r.OpFileloc = col["outputloc"][0].Item

	// get debug
	if col["debug"][0].Item == "1" {
		r.Debug = true
	} else {
		r.Debug = false
	}

	r.Log.Printf("interval: %v\n", r.Interval)
	r.Log.Printf("duration: %v\n", r.Duration)
	r.Log.Printf("cfile location: %v\n", r.CFileLoc)
	r.Log.Printf("output location: %v\n", r.OpFileloc)
	r.Log.Printf("output location: %v\n", r.Debug)

	r.GetNodePodNames()

	return nil
	//fmt.Printf("%s\n%s\n", res.Nodes, res.Pods)
}

func (r *Resources) HandleSignal()  {

	// handle system signals
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, os.Kill, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGSEGV, syscall.SIGFPE)

	f, err := os.OpenFile(r.OpFileloc + strconv.Itoa(os.PathSeparator) + strconv.Itoa(os.Getpid()) + ".pid", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		r.Log.Printf("Create file %s failed!\n", f.Name())
	}

	go func() {
		for {
			select {
			case sig := <- sigs:
				r.Log.Printf("System signal '%v' received! Bye bye!", sig)
				f.Close()
				os.Remove(r.OpFileloc + strconv.Itoa(os.PathSeparator) + strconv.Itoa(os.Getpid()) + ".pid")
				time.Sleep(100 * time.Millisecond)
				os.Exit(2)
			}
		}
	}()

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
//
//func (r *Resources) GetNodePodResource() {
//	config, err := rest.InClusterConfig()
//	if err != nil {
//		panic(err.Error())
//	}
//	// creates the clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
//	for _, item := range pods.Items {
//		if strings.HasPrefix(item.Name, "isdk-ftp-pm-deployment") {
//			r.Pods = append(r.Pods, item.Name)
//		}
//	}
//
//	nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
//	for _, item := range nodes.Items {
//		if strings.HasPrefix(item.Name, "isdk-cluster") {
//			r.Nodes = append(r.Nodes, item.Name)
//		}
//	}
//
//}

//
func (r *Resources) ConvertCmdOutput(kubeOut []byte, name string, str *string) []string  {
	var kubeCmd *exec.Cmd
	if strings.HasPrefix(name, "isdk-cluster") {
		kubeCmd = exec.Command("sh", "-c", "kubectl top node " + name + " | grep isdk")
		//r.Log.Printf("kubectl top node '%s'\n", name)
	} else {
		//kubeCmd = exec.Command("kubectl", "top", "pod", name)
		kubeCmd = exec.Command("sh", "-c", "kubectl top pod " + name + " | grep isdk")
		//r.Log.Printf("kubectl top pod '%s'\n", name)
	}

	kubeOut, err := kubeCmd.Output()
	if err != nil {
		r.Log.Printf("cmd (for %s) failed : %s\n", name, err.Error())
		return []string{}
	} else {
		r.Log.Printf("cmd (for %s) pass!\n", name)
	}

	regexp.Compile("[ ]+")
	//*str = strings.Split(string(kubeOut), "\n")[1]
	*str = string(kubeOut)
	// replace multiple blanks to one
	ReplMB(str, "[ ]+", " ")
	// replace one blank to a comma
	ReplMB(str, " ", ",")
	// drop the last comma of the string
	tmp := *str
	*str = tmp[:len(tmp) - 1]

	// split the string on comma, and replace the unit of each field
	tmpstrs := strings.Split(*str, ",")
	//var joinbackstr string
	if strings.HasPrefix(name, "isdk-cluster") {  // node
		tmpstrs[1] = tmpstrs[1][:len(tmpstrs[1]) - 1]
		tmpstrs[2] = tmpstrs[2][:len(tmpstrs[2]) - 1]
		tmpstrs[3] = tmpstrs[3][:len(tmpstrs[3]) - 2]
		tmpstrs[4] = tmpstrs[4][:len(tmpstrs[4]) - 1]
	} else {   // pod
		tmpstrs[1] = tmpstrs[1][:len(tmpstrs[1]) - 1]
		tmpstrs[2] = tmpstrs[2][:len(tmpstrs[2]) - 2]
	}
	//joinbackstr = strings.Join(tmpstrs, ",")
	//*str = joinbackstr
	//fmt.Println(str)
	return tmpstrs
}

func DialMysql(user, password, hostIP, port, database string) (error, *sql.DB)  {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, hostIP, port, database)
	db, err := sql.Open(dbDriver, dataSource)
	//db, err := sql.Open(DBDriver, User + ":" + Password + "@tcp(" + HostIP + ":" + PORT + ")/" + Database)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}  else {
		fmt.Println("DialMysql successfully")
	}
	return db.Ping(), db
}

func InsertData(db *sql.DB, vals []string)  {
	var name, cpu, cpuPercent, memery, memeryPercent string
	if strings.HasPrefix(vals[0], "isdk-cluster") { // node
		name          = vals[0]
		cpu           = vals[1]
		cpuPercent    = vals[2]
		memery        = vals[3]
		memeryPercent = vals[4]
	} else {  // pod
		name          = vals[0]
		cpu           = vals[1]
		memery        = vals[2]
	}

	//timeStamp := time.Now().UTC().Add(4 * 60 * time.Minute + 10 * time.Minute).Format(time.RFC3339)
	timeStamp := time.Now().UTC().Format(time.RFC3339)


	if strings.HasPrefix(vals[0], "isdk-cluster") {  // node
		insForm, err := db.Prepare("INSERT INTO MonitorNode(Date, Name, CPU_m_cores, CPU_percentage, MEMORY_Mi_bytes, MEMORY_percentage) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		} else {
			//fmt.Println("db sql prepare successfully.")
		}

		_, err = insForm.Exec(timeStamp[:19], name, cpu, cpuPercent, memery, memeryPercent)
		//_, err = insForm.Exec("2019-06-04T12:25:49", "isdk-cluster-control-03", 580, 20, 2586, 90)
		if err != nil {
			panic(err.Error())
		} else {
			//fmt.Println("db sql execution successfully.")
		}
		fmt.Println(fmt.Sprintf("%v | %s, %s, %s, %s, %s", timeStamp[:19], name, cpu, cpuPercent, memery, memeryPercent))
	} else {   // pod
		insForm, err := db.Prepare("INSERT INTO MonitorPod(Date, Name, CPU_m_cores, MEMORY_Mi_bytes) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		} else {
			//fmt.Println("db sql prepare successfully.")
		}

		_, err = insForm.Exec(timeStamp[:19], name, cpu, memery)
		//_, err = insForm.Exec("2019-06-04T12:25:49", "isdk-cluster-control-03", 580, 20, 2586, 90)
		if err != nil {
			panic(err.Error())
		} else {
			//fmt.Println("db sql execution successfully.")
		}
		fmt.Println(fmt.Sprintf("%v | %s, %s, %s", timeStamp[:19], name, cpu, memery))
	}
}

func (r *Resources) CmdOutput(list []string, cmd string)  {
	var db *sql.DB
	var err error

	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
	// creates the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}

	if r.Debug {
		r.Log.Println("in CmdOutput.")
	} else {
		err, db = DialMysql(user, password, hostIP, port, database)
		if err != nil {
			r.Log.Panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}  else {
			//fmt.Println("Database created successfully")
		}

		// use given database
		dbUse := "USE forgrafana"
		_, err = db.Exec(dbUse)
		if err != nil {
			panic(err.Error())
		} else {
			r.Log.Println("DB 'forgrafana' selected successfully!")
		}
	}

	for i, v := range list {
		//r.Wg.Add(1)
		go func(r *Resources, id int, name string) {
			//var outpfExist bool = false
			var f *os.File
			var err error
			t := time.NewTicker(r.Interval)

			//tm := time.NewTimer(r.Duration)
			ctx, _ := context.WithTimeout(context.Background(), r.Duration)

			if r.Debug {
				outputFile := r.OpFileloc + name + ".csv"

				f, err = os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					r.Log.Println( fmt.Sprintf("%d | Create file %s failed!\n", id, f.Name()) )
				}

				if strings.HasPrefix(name, "isdk-cluster") {  //node
					if IsEmptyFile(outputFile) {
						//outpfExist = false
						f.WriteString("TIME,NAME,CPU(m cores),CPU%,MEMORY(Mi bytes),MEMORY%\n")
					} else {
						//outpfExist = true
					}
					//f.WriteString("TIME,NAME,CPU(m cores),CPU%,MEMORY(Mi bytes),MEMORY%\n")

				} else {  // pod
					if IsEmptyFile(outputFile) {
						//outpfExist = false
						f.WriteString("TIME,NAME,CPU(m cores),MEMORY(Mi bytes)\n")
					} else {
						//outpfExist = true
					}
				}
			}

			for {
				select {
				case <-t.C:
					var kubeOut []byte
					var str string
					var vals []string
					//var err error
					if r.Debug {
						kubeOut = []byte(time.Now().Format(time.RFC3339))
						str = string(kubeOut)
						now := time.Now().Format(time.RFC3339)
						now = now[:19]

						vals = r.ConvertCmdOutput(kubeOut, name, &str)

						//pods, err := clientset.CoreV1().Pods("default").Get(v, metav1.GetOptions{})
						//if err != nil {
						//	r.Log.Println(err.Error())
						//}

						//r.Log.Println( fmt.Sprintf("%d | %s,%s", id, now, strings.Join(vals, ",")) )
						if len(vals) != 0 {
							f.WriteString(fmt.Sprintf("%s,%s\n", now, strings.Join(vals, ",")))
						}
						//fmt.Printf("%55v | #%d worker gets top info\n", time.Now(), id)
					} else {
						vals = r.ConvertCmdOutput(kubeOut, name, &str)
						InsertData(db, vals)
					}

				//case <- tm.C:
				case <- ctx.Done():
					r.Log.Println( fmt.Sprintf("time is up, #%d monitor worker ends!", id ) )
					//f.Close()
					r.Wg.Done()
					return
				}
			}
		}(r, i, v)
	}
}

func (r *Resources) MonitorKubeOutput(rType string)  {
	r.Log.Printf("in MonitorKubeOutput type: %s\n", rType)
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


func (r *Resources) DoCommand(cmd string) []string {
	//kubeCmd := exec.Command("sh", "-c", "kubectl get node | grep isdk-cluster | grep -v NotReady | awk -F ' ' '{print $1}'")
	kubeCmd := exec.Command("sh", "-c", cmd)
	kubeOut, err := kubeCmd.Output()
	if err != nil {
		r.Log.Panic("get node failed: " + err.Error())
	}

	str1 := strings.TrimSpace(string(kubeOut))
	str1 = strings.TrimLeft(str1, " ")
	str1 = strings.TrimRight(str1, " ")
	items := strings.Split(str1, "\n")
	//for _, v := range items {
	//	r.Log.Printf("node: '%s'\n", v)
	//}
	return items
}

func (r *Resources) GetNodePodNames()  {
	podCmd := "kubectl get pod -l app=isdk-ftp-pm | grep isdk | awk -F ' ' '{print $1}'"
	r.Pods = append(r.Pods, r.DoCommand(podCmd)...)

	nodeCmd := "kubectl get node | grep -v NotReady | grep isdk | awk -F ' ' '{print $1}'"
	r.Nodes = append(r.Nodes, r.DoCommand(nodeCmd)...)
}
