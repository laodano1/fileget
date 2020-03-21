package main

import (
	"log"
	"sync"
	"time"
)

//func myRoutine(wg *sync.WaitGroup, mp *map[int]string)  {
func myRoutine(wg *sync.WaitGroup, mp *[]int)  {
	defer wg.Done()
	tk := time.Tick(2 * time.Second)
	tm := time.After(20 * time.Second)
	log.Printf("in go routine")
	for {
		select {
		case <- tk:
			log.Printf("mp => %v", mp)
		case <- tm:
			log.Printf("go routine return!")
			return
		}
	}
}

func main() {
	//mp := make(map[int]string)
	mp := make([]int, 0)

	mp = append(mp, 1)
	//for i := 0; i < 4; i++ {
	//	//mp[i] = fmt.Sprintf("i=%v", i)
	//	mp[i] = i
	//}

	log.Printf("in main go routine")
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go myRoutine(wg, &mp)

	time.Sleep(4 * time.Second)
	//mp[20] = 99
	mp = append(mp, 99)
	//time.Sleep(1 * time.Second)
	//delete(mp, 0)

	wg.Wait()
	log.Printf("bye bye!")
}
