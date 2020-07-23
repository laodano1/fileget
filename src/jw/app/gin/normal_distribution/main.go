package main

import (
	"fileget/util"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type tmpRands struct {
	Mp  map[int64]bool
	Arr []int64
}

var (
   port = ":9999"
)

// sample = NormFloat64() * desiredStdDev + desiredMean
func GetNormFloat64(dev, mean float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.NormFloat64() * dev + mean
}

func GetNormInt64(dev, mean float64) (val int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tmp := r.NormFloat64()
	if tmp > 0 {
		if tmp > 1 {
			//util.Lg.Debugf("tmp: %v, tmp * dev: %v, mean: %v", tmp, tmp-1, int64(mean))
			return
		} else {
			//util.Lg.Debugf("tmp: %v, tmp * dev: %v, mean: %v", tmp, 1-tmp, int64(mean))
		}
	} else {
		if tmp < -1 {
			return
		}
		//util.Lg.Debugf("tmp: %v, tmp * dev: %v, mean: %v", tmp, tmp*-1, int64(mean))
	}
	//util.Lg.Debugf("tmp * dev: %v, mean: %v", tmp%1, int64(mean))
	return int64(tmp * dev + mean)
}

func GetVal(tR *tmpRands) {
	for i := 0; i < 100; i++ {
		item := GetNormInt64(500, 1500)
		if item == 0 {continue}
		//util.Lg.Debugf("rand: %v", item)
		tR.Mp[item] = true
		tR.Arr = append(tR.Arr, item)
	}
}

func GenerateVal(tR *tmpRands) {
	//tR := &tmpRands{}
	tR.Mp = make(map[int64]bool)
	for  {
		GetVal(tR)
		if len(tR.Arr) != 0 {break}
	}
	if tR.Arr[int64(len(tR.Arr)/2)] < 1000 || tR.Arr[int64(len(tR.Arr)/2)] > 2000 {
		util.Lg.Debugf("********************** bad value!!!")
	}
}

func StartWebServer() {
	gin.ForceConsoleColor()
	e := gin.Default()


	e.GET("/", func(c *gin.Context) {
		tR := &tmpRands{}
		//tR.Mp = make(map[int64]bool)
		//for  {
		//	GetVal(tR)
		//	if len(tR.Arr) != 0 {break}
		//}
		//if tR.Arr[int64(len(tR.Arr)/2)] < 1000 || tR.Arr[int64(len(tR.Arr)/2)] > 2000 {
		//	util.Lg.Debugf("********************** bad value!!!",)
		//}
		GenerateVal(tR)
		c.IndentedJSON(http.StatusOK, tR.Arr[int64(len(tR.Arr)/2)])
	})

	if err := e.Run(port); err != nil {
		util.Lg.Errorf("error: %v", err)
	}

}

func main() {
	//GetNormFloat64(1000, 15000)

	//for i := 0; i < b.N; i++ {
	//	util.Lg.Debugf("rand: %v", GetNormInt64(1000, 15000))
	//}
	StartWebServer()

}
