package main

import (
	"fmt"
	"time"
)

//type worker struct {
//
//}

type pool struct {
	WorkNum    int
	JobChan    chan task
	//Workers    map[int]*worker
	OffDuty    chan bool
	RollCall   chan bool
	Done       chan bool
}

//func newWorker() (w *worker) {
//	return &worker{}
//}

//func (p *pool) intiWorker() {
//	go func() {
//		select {
//		case <- :
//
//		}
//	}()
//}

func newPool() (p *pool) {
	return new(pool)
}


func (p *pool) initPool(workNum int) {
	logger.Debugf("init pool")
	p.WorkNum  = workNum
	p.OffDuty  = make(chan bool, workNum)
	p.RollCall = make(chan bool, workNum)
	p.Done = make(chan bool)
	for i := 0; i < workNum; i++ {
		logger.Debugf("init worker: %v", i)
		go func(id int) {
			for {
				select {
				case taskItem := <- p.JobChan:
					logger.Debugf("worker(%v) execute task", id)
					taskItem.t()
				case <- p.OffDuty:
					logger.Debugf("worker(%v) is off duty. Bye bye!", id)
					p.RollCall <- true
					return
				}
			}
		}(i)
	}
}

//func (p *pool) setWorkTime(t int) {
//	go func() {
//		time.AfterFunc(10 * time.Second, func() {
//			p.OffDuty <- true
//		})
//	}()
//}

func (p *pool) isFullRollCall() {
	go func() {
		logger.Debugf("roll call check")
		num := 0
		for {
			select {
			case <- p.RollCall:
				num++
				if num == p.WorkNum {
					p.Done <- true
				} else {
					//logger.Debugf("num: %v", num)
				}
			}
		}
	}()
}

func (p *pool) takeWork(intiTaskNum int) {
	tk1 := time.Tick(2 * time.Second)
	tk2 := time.Tick(5 * time.Second)

	go func() {
		logger.Debugf("pool takeWork")
		for i := 0; i < intiTaskNum; i++ {
			fmt.Printf("%v ", i)
			t := task{
				t:func() {
					logger.Debugf("this is a task 0! %v", time.Now().Format(time.RFC3339Nano))
				}}
			p.JobChan <- t
		}
		fmt.Printf("\n")
		logger.Debugf("pool takeWork 2")
		for {
			select {
			case <- tk1:
				logger.Debugf("generate tk1")
				t := task{
					t:func() {
						logger.Debugf("this is a task 1! %v", time.Now().Format(time.RFC3339Nano))
					}}
				p.JobChan <- t
			case <- tk2:
				logger.Debugf("generate tk2")
				t := task{
					t:func() {
						logger.Debugf("this is a task 2! %v", time.Now().Format(time.RFC3339Nano))
					}}
				p.JobChan <- t
			}

		}
	}()

}







