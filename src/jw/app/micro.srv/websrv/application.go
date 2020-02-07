package main

import (
	"github.com/micro/go-micro/web"
)

type MyService struct {
	//web service
	myWebService web.Service
}

type DBOperation interface {
	Add(cnt interface{}) error
	Del(key interface{}) error
	Update(key interface{}) error
	Select(key interface{}) ([]interface{}, error)
}
