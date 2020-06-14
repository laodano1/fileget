package main

import (
	"context"
	"fileget/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	var err error

	clientOpts := options.Client().ApplyURI("mongodb://10.0.0.130:27017,10.0.0.130:27018,10.0.0.130:27019/jwtestdb?replicaSet=my-mongo-set")
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		util.Lg.Errorf("Connect mongodb replset failed: %v", err)
		return
	}

	res, er := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		util.Lg.Errorf("ListDatabases failed: %v", er)
		return
	}

	util.Lg.Debugf("%v", res.Databases)
	type myData struct {
		Id   int64  `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
		Age  int32  `json:"age" bson:"age"`
	}
	md := &myData{
		Id:   3,
		Name: "jw",
		Age:  30,
	}
	//client.Database("jwtestdb").Collection("jwcol").InsertOne(ctx, bson.M{"hello": "world"})
	result, err := client.Database("jwtestdb").Collection("jwcol").InsertOne(ctx, md)
	if err != nil {
		util.Lg.Errorf("InsertOne failed: %v", er)
		return
	}
	util.Lg.Debugf("result: %v", result)

}
