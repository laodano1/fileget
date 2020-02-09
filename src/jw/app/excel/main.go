package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/tealeg/xlsx"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type myRow struct {
	Cells []string `json:"row"`
}

type mySheet struct {
	Name string  `json:"name"`
	Rows []myRow `json:"sheet"`
}

type myExcel struct {
	Name       string
	SheetNames []string
	Sheets     []mySheet `json:"sheets"`
}

var tmplStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Title }}</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

</head>
<body>
	<!--img style="width:50%;height:100px;" src="https://cn.bing.com/th?id=OHR.SneezeSpring_EN-CN8669316656_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp" class="img-fluid" alt="cover image">
    < table-striped table-responsive-sm bg-info bg-success -->
	
	<ul class="nav nav-tabs" id="myTab" role="tablist">
		{{range $idx, $name := .SheetNames}}
	  <li class="nav-item">
		{{if eq $idx 0 }}
		<a class="nav-link active" data-toggle="tab"  href="#tab-pane-{{$idx}}" role="tab">{{ . }}</a>
        {{else}}
		<a class="nav-link" data-toggle="tab"  href="#tab-pane-{{$idx}}" role="tab" >{{ . }}</a>
		{{end}}
	  </li>
		{{end}}
		<div class="dropdown">
		  <button class="btn btn-info dropdown-toggle" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
			Page: {{ .Title }}
		  </button>
		  <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
			{{range $name := .ExcelNames}}
			{{if ne $name ""}}
			<a class="dropdown-item" href="/{{ SplitFileName $name}}">{{ SplitFileName $name}}</a>
			{{end}}
			{{end}}
		  </div>
		</div>
	</ul>
	<div class="tab-content">
		{{$myExcelCnt := .Sheets}}
		{{range $idx, $name := .SheetNames}}
		{{if eq $idx 0}}
		<div id="tab-pane-{{$idx}}" class="tab-pane active" role="tabpanel">
		{{else}}
		<div id="tab-pane-{{$idx}}" class="tab-pane" role="tabpanel">
		{{end}}
			<table class="table table-bordered table-hover table-responsive-sm ">
			  <thead class="thead-light">
				<tr>
					{{range (index (index $myExcelCnt.Sheets $idx).Rows 0).Cells}}
					<th style='text-align: center;' scope="col">
						{{ . }}
					</th>
					{{end}}
				</tr>
			  </thead>
			  <tbody>
				{{range $idx, $row := (index $myExcelCnt.Sheets $idx).Rows}}
				{{if ne $idx 0}}
					<tr>
						{{range $cell := $row.Cells}}
							{{if ne $cell ""}}
							<td style='text-align: center;'>
								{{ $cell }}
							</td>
							{{end}}
						{{end}}
					</tr>
				{{end}}
				{{end}}
			  </tbody>
			</table>
		</div>
		{{end}}
	</div>
    <script src="https://code.jquery.com/jquery-3.4.1.slim.min.js" integrity="sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js" integrity="sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js" integrity="sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6" crossorigin="anonymous"></script>
	
	<script>
		$(function () {
			$('#myTab li:first-child a').tab('show')
		  })
	</script>

</body>
</html>
`

type allExcels struct {
	cnt map[string]*myExcel //string: file name
}

var (
	excelFiles []string
	//excelFile  = "mytest.xlsx"
	//me          *myExcel
	logger     = golog.New("myExcel")
	sheetNames []string

	excelSuffix = ".xlsx"
)

func HasExist(allItems []string, item string) bool {
	var itemExist bool
	for _, name := range allItems {
		if name == item {
			itemExist = true
			return itemExist
		}
	}

	return itemExist
}

func (ae *allExcels) ExtractExcelFile(addPrefix bool) (err error) {
	if len(ae.cnt) < len(excelFiles) {
		// 有excel文件被删除了
		logger.Infof("有excel文件被删除了")
		excelFiles = make([]string, 0)
		for _, ce := range ae.cnt {
			excelFiles = append(excelFiles, ce.Name)
		}
	}

	for path, me := range ae.cnt {
		//me = new(myExcel)
		xlFile, err := xlsx.OpenFile(path)
		if err != nil {
			logger.Errorf("open excel file error: %v", err)
			return err
		}

		me.Sheets = make([]mySheet, 0)
		me.Name = path
		me.SheetNames = make([]string, 0)

		if !HasExist(excelFiles, path) {
			excelFiles = append(excelFiles, path)
		} else {
			continue
		}

		for _, sheet := range xlFile.Sheets {
			logger.Infof("excel: %-14v, sheet name: %-9v", path, sheet.Name)
			me.SheetNames = append(me.SheetNames, sheet.Name)
			ms := mySheet{}
			ms.Name = sheet.Name
			ms.Rows = make([]myRow, 0)
			for r, row := range sheet.Rows {
				mr := myRow{}
				mr.Cells = make([]string, 0)
				for i, cell := range row.Cells {
					if r != 0 {
						if cell.IsTime() {
							cell.SetFormat("yyyy-mm-dd hh:mm:ss")
							v, err := cell.FormattedValue()
							if err != nil {
								logger.Errorf("format value failed: %v", err)
								return err
							}
							if addPrefix {
								mr.Cells = append(mr.Cells, fmt.Sprintf("%v:%v", sheet.Rows[0].Cells[i], v))
							} else {
								mr.Cells = append(mr.Cells, fmt.Sprintf("%v", v))
							}

							//logger.Infof("%s", v)
						} else {
							//logger.Infof("%s", cell.Value)
							if addPrefix {
								mr.Cells = append(mr.Cells, fmt.Sprintf("%v:%v", sheet.Rows[0].Cells[i], cell.Value))
							} else {
								mr.Cells = append(mr.Cells, fmt.Sprintf("%v", cell.Value))
							}
						}
					} else {
						mr.Cells = append(mr.Cells, cell.Value)
					}
				}
				ms.Rows = append(ms.Rows, mr)
			}
			me.Sheets = append(me.Sheets, ms)
		}
	}

	return
}

func (me *myExcel) pageHandle(ctx *gin.Context) {
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"SplitFileName": func(fullStr string) (subStr string) {
			return strings.Split(fullStr, ".")[0]
		},
	}).Parse(tmplStr)) // Create a template
	dt := &gin.H{
		"SheetNames": me.SheetNames,
		"Title":      strings.Split(me.Name, ".")[0],
		"ExcelNames": excelFiles,
		"Sheets":     me,
	}

	htmlRender := render.HTML{
		Template: tmpl,
		Name:     "",
		Data:     dt,
	}
	ctx.Render(http.StatusOK, htmlRender)
}

func init() {

}

func (ae *allExcels) GetAllExcelFiles() {

	fileInfors, err := ioutil.ReadDir(".")
	if err != nil {
		logger.Errorf("cannot read dir, error: %v", err)
		panic(err)
	}
	ae.cnt = make(map[string]*myExcel)
	for _, fileInfo := range fileInfors {
		if strings.HasSuffix(fileInfo.Name(), excelSuffix) {
			ae.cnt[fileInfo.Name()] = new(myExcel)
		}
	}
}

func main() {
	var addPrefix bool
	var reload bool
	flag.BoolVar(&addPrefix, "ap", false, "if add prefix of cell value")
	flag.BoolVar(&reload, "reload", false, "if reload")

	flag.Parse()
	var AllExcels allExcels
	//AllExcels.cnt = make(map[string]*myExcel)

	var err error
	AllExcels.GetAllExcelFiles()
	err = AllExcels.ExtractExcelFile(addPrefix)
	if err != nil {
		logger.Errorf("extract excel file failed: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"result": false, "error": "Method Not Allowed"})
		return
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"result": false, "error": "Endpoint Not Found"})
		return
	})

	allPaths := make(map[string]bool)
	for fileName, e := range AllExcels.cnt {
		pathStr := fmt.Sprintf("/%v", strings.Split(fileName, ".")[0])
		router.GET(pathStr, e.pageHandle)
		allPaths[pathStr] = true
		logger.Infof("gin handle in http path: %v", pathStr)
	}

	go func() {
		tk := time.Tick(20 * time.Second)
		for {
			select {
			case <-tk:
				AllExcels.GetAllExcelFiles()
				err = AllExcels.ExtractExcelFile(addPrefix)
				if err != nil {
					logger.Errorf("extract excel file failed: %v", err)
				}

				for fileName, e := range AllExcels.cnt {
					pathStr := fmt.Sprintf("/%v", strings.Split(fileName, ".")[0])
					if _, ok := allPaths[pathStr]; !ok { // add new http path
						router.GET(pathStr, e.pageHandle)
						logger.Infof("gin handle in http path: %v", pathStr)
						allPaths[pathStr] = true
					}
				}
			}
		}

	}()

	logger.Infof("start to listen on :9999")
	if err := router.Run(":9999"); err != nil {
		logger.Errorf("gin router run failed: %v", err)
	}

}
