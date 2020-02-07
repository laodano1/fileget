package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type webServer struct {
	name string
	srv  *gin.Engine
}

var WebServer *webServer

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
	c.String(http.StatusOK, "home page | "+time.Now().Format(time.RFC3339Nano))
}

func aboutpage(c *gin.Context) {
	c.String(http.StatusOK, "about page | "+time.Now().Format(time.RFC3339Nano))
}

func (ws *webServer) Start() {
	ws.name = "my-gin-web"
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

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

	ws.srv = router
}

// reload config file
func (ws *webServer) Reload() {

}

func init() {
	WebServer = new(webServer)
	WebServer.Start()
}
