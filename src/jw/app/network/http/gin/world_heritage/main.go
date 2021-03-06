package main

func main() {
	if err := PreStartActions(); err != nil {
		return
	}
	ms := NewGinServer()
	ms.AddRoutes()
	ms.Start()
}

