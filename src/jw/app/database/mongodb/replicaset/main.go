package main

import (
	"fileget/util"
	"gopkg.in/mgo.v2"
	"time"
)

func main() {
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	//session, err := mgo.Dial("mongodb://10.0.0.130:27017,10.0.0.130:27018,10.0.0.130:27019/jwtestdb?replicaSet=my-mongo-set")
	//if err != nil {
	//	util.Lg.Errorf("dial mongodb error: %v", err)
	//	return
	//}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{
			"10.0.0.130:27017",
			"10.0.0.130:27018",
			"10.0.0.130:27019",
		},
		Timeout:        10 * time.Second,
		Database:       "jwtestdb",
		ReplicaSetName: "my-mongo-set",
	})
	if err != nil {
		util.Lg.Errorf("dial mongodb error: %v", err)
		return
	}

	util.Lg.Debugf("num: %v,  servers: %v", len(session.LiveServers()), session.LiveServers())
	n, _ := session.DB("jwtestdb").C("jwcol").Count()
	util.Lg.Debugf("jwcol count: %v", n)

}
