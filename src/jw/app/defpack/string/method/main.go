package main

import (
	"fileget/util"
	"fmt"
)
var (
	v1 string
)

type str = string

//
// Cannot define new methods on non-local type 'builtin.string'
//
//func (s str) why() string {
//	return "why"
//}

type Person struct {
	age int
}
func (p Person) howOld() int {
	return p.age
}
func (p *Person) growUp() {
	p.age += 1
}
//func main() {
//	// qcrao 是值类型
//	qcrao := Person{age: 18}
//	// 值类型 调用接收者也是值类型的方法
//	fmt.Println(qcrao.howOld())
//
//	// 值类型 调用接收者是指针类型的方法
//	qcrao.growUp()
//	fmt.Println(qcrao.howOld())
//	// ----------------------
//
//	// stefno 是指针类型
//	stefno := &Person{age: 100}
//	// 指针类型 调用接收者是值类型的方法
//	fmt.Println(stefno.howOld())
//	// 指针类型 调用接收者也是指针类型的方法
//	stefno.growUp()
//	fmt.Println(stefno.howOld())
//}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {   // 实现了接受者是值类型的方法，就自动实现了指针类型的方法
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() { // 实现了接受者是指针类型的方法，不会自动实现值类型的方法
	fmt.Printf("I am debuging %s language\n", p.language)
}
//
//func main() {
//	var c coder = &Gopher{"Go"}
//	//var c coder = Gopher{"Go"}
//	c.code()
//	c.debug()
//}


type MyError struct {}

func (i MyError) Error() string {
	return "MyError"
}
//
//func main() {
//	err := Process()
//	fmt.Println(err)
//
//	fmt.Println(err == nil)
//}

func Process() error {
	var err *MyError = nil
	return err
}



type myWriter struct {

}

func (w myWriter) Write(p []byte) (n int, err error) {
	return
}

//func (w *myWriter) Write(p []byte) (n int, err error) {
//	return
//}

//func main() {
//	// 检查 *myWriter 类型是否实现了 io.Writer 接口
//	var _ io.Writer = (*myWriter)(nil)
//
//	// 检查 myWriter 类型是否实现了 io.Writer 接口
//	var _ io.Writer = myWriter{}
//}

type myString struct {
	Id   int32
	Name string
}

// 现实String()接口，则在打印的时候默认会调用它
func (ms myString) String() string {
	return fmt.Sprintf("{id: %v, name: %v}", ms.Id, ms.Name)
}

func main() {
	ms := myString{
		Id: 1,
		Name: "jw",
	}

	util.Lg.Debugf("myString: %v", ms)

}

//func main() {
//	var i int = 9
//
//	var f float64
//	f = float64(i)
//	fmt.Printf("%T, %v\n", f, f)
//
//	f = 10.8
//	a := int(f)
//	fmt.Printf("%T, %v\n", a, a)
//
//	//s := []int(i)
//	//s := []int{i}
//}