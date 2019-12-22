package main

import (
	"github.com/gin-gonic/gin"
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

func homepage(c *gin.Context) {
	c.String(http.StatusOK, "hellllll worlllld")
}

func aboutpage(c *gin.Context) {
	c.String(http.StatusOK, "about page")
}

func main() {

	webSrv := web.NewService(
		web.Name("go.micro.web.form"),
		func(o *web.Options) {
			o.Address = ":9999"
		},
	)

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
	router.GET("/about", aboutpage)

	webSrv.Handle("/", router)
	if err := webSrv.Run(); err != nil {
		log.Fatal("go micro web server error: ", err)
	}

}
