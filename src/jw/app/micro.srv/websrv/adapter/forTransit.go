package main

type msgTransit struct {
	msgCh chan interface{} //

}

func NewMsgTransit() *msgTransit {
	Ch := make(chan interface{})
	//go func() {
	//	for {
	//		select {
	//		case item := <-Ch:
	//			lg.Debugf("get msg: %v", item)
	//
	//		}
	//	}
	//}()
	return &msgTransit{msgCh: Ch}
}

func (mt *msgTransit) AddMsg(msg interface{}) {
	mt.msgCh <- msg
}

//
//func (mt *msgTransit) operateMsg()  {
//
//}
