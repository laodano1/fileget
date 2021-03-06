package main

import "fmt"

type dmap map[int]int

type pmap map[int]dmap

func main() {
	pm := make(pmap)

	dm := make(dmap)
	dm[0] = 0
	dm[1] = 1

	pm[0] = dm


	fmt.Printf("num: %v\n", len(pm))
	fmt.Printf("num: %v\n", len(dm))

}
