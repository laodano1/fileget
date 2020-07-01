package main

import "fmt"

// 代码函数里面的代码块
func main() {
	value := 1
	//fmt.Println(value) // prints 1
	fmt.Printf("1 value: %v, addr of value: %p\n", value, &value) // prints 2
	{
		fmt.Printf("2 value: %v, addr of value: %p\n", value, &value) // prints 2
		//fmt.Println(value)                                          // prints 1
		value = 2
		//value := 2
		//fmt.Printf("value: %v, addr of value: %p\n", value, &value) // prints 2
	}
	//fmt.Println(value)     // prints 1 (bad if you need 2)
	fmt.Printf("4 value: %v, addr of value: %p\n", value, &value) // prints 2
}
