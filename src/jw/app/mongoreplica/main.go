package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)
func main() {
	// user, password, database, rs0分别为用户名、密码、数据库、副本集，请自行修改
	URL := "mongodb://10.0.0.57:27017,10.0.0.57:27018,10.0.0.57:27019/?replicaSet=rep_Wind"
	session, err := mgo.Dial(URL)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	session.SetMode(mgo.Monotonic, true)
	Conn := session.DB("windplatform")
	// 获得 Mongodb的连接后，再就可以进行各种CRUD啦

	names, err := Conn.CollectionNames()
	if err != nil {
		panic(err)
	}

	log.Printf("collections: %v", names)
}

