package main

import (
	"fmt"
	"log"
	"reflect"
)

type (
	People struct {
	}

	Teacher struct {
		People
	}
)

func (p *People) forA() {
	log.Printf("people a\n")
	p.forB()
}

func (p *People) forB() {
	log.Printf("people b\n")
}

func (t *Teacher) forB() {
	log.Printf("teacher b")
}

const (
	a = iota
	b
	c = "zz"
	d
	f = iota
)

func main1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
	//list := new([]int)
	////i := 1
	//*list = append(*list, 1)
	//
	//sn1 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//
	//sn2 := struct {
	//	age  int
	//	name string
	//}{age: 11, name: "qq"}
	//
	//if sn1 == sn2 {
	//	fmt.Println("sn1 == sn2")
	//}

	//println(DeferFunc1(1))
	//println(DeferFunc2(1))
	//println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
