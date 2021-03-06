package main

import (
	"github.com/davyxu/golog"
	"github.com/jinzhu/gorm"
)

type myDBClient struct {
	myDB *gorm.DB
}

var MyDBClient myDBClient
var logger golog.Logger
var host string = "10.0.0.146:3306"

func initMysql(host string) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", host)
	if err != nil {
		logger.Errorf("gorm open failed: %v", err)
	}

	return
}

func init() {
	db, err := initMysql(host)
	if err != nil {
		logger.Errorf("init mysql connection failed: %v", err)
	}

	MyDBClient.myDB = db
}
