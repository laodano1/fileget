package main

type result struct {
	val interface{}
	err error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key       string
	response  chan<- result
}

type server struct {
	requests chan request
}

type Func func(key string) (interface{}, error)

func NewServer(f Func) *server {
	s := &server{
		requests: make(chan request),
	}
	go s.Start(f)
	return s
}

func (e *entry) call(f Func, key string)  {
	e.res.val, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(rsp chan<- result)   {
	<- e.ready
	rsp <- e.res
}

func (s *server) Start(f Func)  {
	cache := make(map[string]*entry)
	for req := range s.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (s *server) Get(key string) (ret interface{}, err error) {
	rsp := make(chan result)
	s.requests <- request{response: rsp}
	res := <- rsp
	return res.val, res.err
}

func main() {



}
