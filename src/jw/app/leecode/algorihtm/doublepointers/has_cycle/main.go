package main
//
//import (
//	"container/list"
//	"fileget/util"
//)
//
//func has_cycle(head *list.Element) bool {
//	if head.Next() == nil { return false}
//
//	l1 := head
//	l2 := head.Next()
//	for l1.Next() != nil && l2 != nil && l2.Next() != nil {
//		if l1.Next() == l2 {
//			return true
//		}
//	}
//
//	return false
//}
//
//func main() {
//	l := list.New()
//	l.PushBack(1)
//	l.PushBack(2)
//	l.PushBack(3)
//	util.Lg.Debugf("has cycle: %v", has_cycle(l.Front()))
//}
