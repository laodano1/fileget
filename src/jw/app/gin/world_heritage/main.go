package main

func main() {
	ms := NewGinServer()
	ms.AddRoutes()
	ms.Start()
}

