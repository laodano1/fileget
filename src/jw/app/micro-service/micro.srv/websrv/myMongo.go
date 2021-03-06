package main

import (
	"gopkg.in/mgo.v2"
	"time"
)

type myMongoClient struct {
	myDB *mgo.Database
}

var MyMongoDB myMongoClient
var mgHost string = "10.0.0.146"
var dbName string = "myMgo"

func initMongo(host string) (db *mgo.Database, err error) {
	ss, err := mgo.DialWithTimeout("mongodb://"+host, 3*time.Second)
	if err != nil {
		logger.Errorf("mgo dial failed: %v", err)
	}

	db = ss.DB(dbName)

	return
}

func init() {
	db, err := initMongo(mgHost)
	if err != nil {
		logger.Errorf("init mongo failed: %v", err)
	}

	MyMongoDB.myDB = db

}
