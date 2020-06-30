package main

import "fileget/util"

type S struct {
	Name string
	slc  []int
	arr  [3]int
	num  int
}

func main() {
	myArr := [3]int{0, 0, 0}
	mp := make(map[int]S)   // *********** value为结构体类型
	mp[0] = S{Name: "haha", slc: make([]int, 0), arr: myArr}
	util.Lg.Debugf("name: %v", mp[0].Name)

	//mp[0].arr[0] = 1        // 值类型不能赋值，但是引用类型可以赋值
	//mp[0].Name = "update"   // : cannot assign to struct field mp[0].Name in map
	//mp[0].num = 1           // : cannot assign to struct field mp[0].num in map
	mp[0].slc[0] = 1          // 可以给slice赋值  值类型不能赋值，但是引用类型可以赋值
	//错误原因:
	// 1) x = y 这种赋值的方式，你必须知道 x的地址，然后才能把值 y 赋给 x。
	// 2) 但 go 中的 map 的 value 本身是不可寻址的，因为 map 的扩容的时候，可能要做 key/val pair迁移,所以拒绝直接寻址，避免产生野指针；
	// 3) value 本身地址是会改变的

	// 4) 不支持寻址的话又怎么能赋值呢
	mp1 := make(map[int]*S)  // ********** value修改为结构体指针类型
	mp1[0] = &S{Name: "old"}
	mp1[0].Name = "new"
	// 1) 由刚刚得推断我们可以发现，只要知道了被修改值的地址，我们就可以修改它了
	// 2) 所以我们使用指针和引用保证每次赋值都可以找到地址
	// 3) 就可以实现 map 的结构体赋值了

	mp1[1] = &S{Name: "2", arr: make([]int, 0)}
	mp1[1].arr[0] = 1

}
