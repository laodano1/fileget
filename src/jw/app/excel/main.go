package main

import (
	"flag"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/tealeg/xlsx"
	"html/template"
	"net/http"
	"time"
)

type myRow struct {
	Cells []string `json:"row"`
}

type mySheet struct {
	Rows []myRow `json:"sheet"`
}

type myExcel struct {
	Sheets []mySheet `json:"sheets"`
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
	
	<ul class="nav nav-tabs tab-content">
		{{range $idx, $name := .SheetNames}}
	  <li class="nav-item">
		{{if eq $idx 0 }}
		<a class="nav-link active" href="#">{{ . }}</a>
        {{else}}
		<a class="nav-link" href="#">{{ . }}</a>
		{{end}}
	  </li>
		{{end}}
	</ul>
	<div class="tab-content">
		{{$myExcelCnt := .Sheets}}
		{{range $idx, $name := .SheetNames}}
		<div class="tab-pane active" id="home" role="tabpanel" aria-labelledby="home-tab">
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
</body>
</html>
`

var (
	excelFile  = "mytest.xlsx"
	me         *myExcel
	logger     = golog.New("myExcel")
	sheetNames []string
)

func ExtractExcelFile(path string, addPrefix bool) (me *myExcel, err error) {
	//var xlFile *xlsx.File
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		logger.Errorf("open excel file error: %v", err)
		return nil, err
	}

	me = &myExcel{}
	me.Sheets = make([]mySheet, 0)
	sheetNames = make([]string, 0)

	for _, sheet := range xlFile.Sheets {
		logger.Infof("sheet name: '%v'", sheet.Name)
		sheetNames = append(sheetNames, sheet.Name)
		ms := mySheet{}
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
							return nil, err
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

	return
}

func homepage(ctx *gin.Context) {
	tmpl := template.Must(template.New("").Parse(tmplStr)) // Create a template
	dt := &gin.H{
		"SheetNames": sheetNames,
		"Title":      excelFile,
		"Sheets":     me,
		//"THead":      me.Sheets[0].Rows[0].Cells,
		//"TBody":      me.Sheets[0].Rows[1:],
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

func main() {
	var addPrefix bool
	var reload bool
	flag.BoolVar(&addPrefix, "ap", false, "if add prefix of cell value")
	flag.BoolVar(&reload, "reload", false, "if reload")

	flag.Parse()

	var err error
	me, err = ExtractExcelFile(excelFile, false)
	if err != nil {
		logger.Errorf("extract excel file failed: %v", err)
	}

	go func() {
		tk := time.Tick(20 * time.Second)
		for {
			select {
			case <-tk:
				me, err = ExtractExcelFile(excelFile, addPrefix)
				if err != nil {
					logger.Errorf("extract excel file failed: %v", err)
				}
			}
		}

	}()

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

	router.GET("/", homepage)

	logger.Infof("start to listen on :9999")
	if err := router.Run(":9999"); err != nil {
		logger.Errorf("gin router run failed: %v", err)
	}

}
