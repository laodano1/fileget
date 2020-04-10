package main

import (
	"fmt"
	"github.com/davyxu/golog"
	"github.com/gin-gonic/gin"
	"net/http"
	"unsafe"
)

type Person struct {
	//ID   string `uri:"id"   binding:"required,uuid"`
	ID   string `uri:"id" `
	Name string `uri:"name" binding:"required"`
}

func main() {
	x := make([]int, 1, 2)
	y := x
	fmt.Printf("%p <=> %p\n", x, y)
	y = append(y, 1)
	fmt.Println(x, y)
	x = append(x, 2)
	fmt.Printf("%p <=> %p\n", x, y)
	fmt.Println(x, y)
	x = append(x, 3)
	fmt.Printf("%p <=> %p\n", x, y)
	fmt.Println(x, y)

	var s string = "abcde"
	var c = 'A'
	fmt.Printf("%d\n", c)
	var arr = [4]int{}
	type st struct {
		i int
	}
	fmt.Printf("%d, %d, %d\n", unsafe.Sizeof(x), unsafe.Sizeof(y), unsafe.Sizeof(st{}))
	fmt.Printf("%d, %d, %d\n", unsafe.Sizeof(s), unsafe.Sizeof(arr), unsafe.Sizeof(c))
}

func main111() {
	lg := golog.New("mylist")

	gin.ForceConsoleColor()
	r := gin.Default()

	// TODO: jsonp
	//r.GET("/JSONP?callback=x", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"foo": "bar",
	//	}
	//
	//	//callback is x
	//	// Will output  :   x({\"foo\":\"bar\"})
	//	c.JSONP(http.StatusOK, data)
	//})

	// TODO: binding uri
	//r.GET("/:name/:id", func(c *gin.Context) {
	//	var person Person
	//	if err := c.ShouldBindUri(&person); err != nil {
	//		c.JSON(400, gin.H{"msg": err})
	//		return
	//	}
	//	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	//})

	r.Use(gin.Recovery())

	// TODO: ascii json
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})




	// Listen and serve on 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		lg.Errorf("gin run failed: %v", err)
	}
}

