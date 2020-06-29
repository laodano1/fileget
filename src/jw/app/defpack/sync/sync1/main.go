package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			fmt.Printf("i: %v\n", idx)
			m[idx] = idx
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}

