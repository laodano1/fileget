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

func GenerateVal(tR *tmpRands) {
	tR.Mp = make(map[int64]bool)

	min := 2500000
	max := 3000000
	dev  := float64((max - min)/2)
	mean := float64((min + max)/2)
	util.Lg.Debugf("dev: %v, mean: %f", dev, mean)
	GetVal(tR, dev, mean)
}


func GetNormInt64(dev, mean float64) (val int64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//util.Lg.Debugf("%v", r.NormFloat64())
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

func GetVal(tR *tmpRands, dev, mean float64) {
	for {
		item := GetNormInt64(dev, mean)
		if item == 0 {continue}
		//util.Lg.Debugf("rand: %v", item)
		tR.Mp[item] = true
		tR.Arr = append(tR.Arr, item)
		break
	}
}

func GetVal2(dev, mean float64) (val int64) {
	for {
		val = GetNormInt64(dev, mean)
		if val == 0 {continue} else {val = val - val % 10000; break}
	}
	return
}

func GenerateVal2(min, max float64) int64 {
	dev  := (max - min)/2
	mean := (min + max)/2
	//util.Lg.Debugf("dev: %v, mean: %f", dev, mean)
	return GetVal2(dev, mean)
}

func StartWebServer() {
	gin.ForceConsoleColor()
	e := gin.Default()

	e.GET("/", func(c *gin.Context) {

		min := float64(850000)
		max := float64(1650000)
		arr := make([]int64, 0)
		for i := 0; i < 1; i++ {
			arr = append(arr, GenerateVal2(min, max))
		}
		c.IndentedJSON(http.StatusOK, arr)
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
