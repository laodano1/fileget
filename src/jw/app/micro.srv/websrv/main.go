package main

import (
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, `<html>
	//<body>
	//<h1>This is a regular form</h1>
	//<form action="http://localhost:8080/form/submit" method="POST">
	//<input type="text" id="thing" name="thing" />
	//<button>submit</button>
	//</form>
	//<h1>This is a multipart form</h1>
	//<form action="http://localhost:8080/form/multipart" method="POST" enctype="multipart/form-data">
	//<input type="text" id="thing" name="thing" />
	//<button>submit</button>
	//</form>
	//</body>
	//</html>
	//`)
	w.Write([]byte("hello world"))
}

func main() {

	webSrv := web.NewService(
		web.Name("go.micro.web.form"),
		func(o *web.Options) {
			o.Address = ":9999"
		},
	)
	webSrv.HandleFunc("/", index)

	if err := webSrv.Run(); err != nil {
		log.Fatal("go micro web server error: ", err)
	}

}
