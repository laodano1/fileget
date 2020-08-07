package main

import (
	"context"
	"fileget/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"time"
)

type Database struct {
	client *mongo.Client
	db *mongo.Database
	playerCol *mongo.Collection
	questCol *mongo.Collection
}

type robotUnit struct {
	Id   int32  `bson:"_id" `
	Type int32  `bson:"type"`
	V0   int32  `bson:"v0"`
	V1   int32  `bson:"v1"`
	V2   int32  `bson:"v2"`
	Prob int32  `bson:"Prob"`
}

var instance *Database

func InitDatabase(uri string, dbName string) (err error) {
	db := Database{}
	db.client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = db.client.Connect(ctx)
	if err != nil {
		return err
	}
	db.db = db.client.Database(dbName)
	instance = &db
	return nil
}

func GetDatabase() *Database {
	return instance
}

func (this *Database) GetDBCli() *mongo.Client {
	return instance.client
}

func RandIntRange(left int, right int) int {
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(right-left+1) + left
}

func GetUnit() {
	urlStr, dbname := "mongodb://10.0.0.252:27017,10.0.0.251:27017/?replicaSet=xxx", "meta"
	InitDatabase(urlStr, dbname)
	cur, err := instance.GetDBCli().Database("meta").Collection("Unit").Find(context.TODO(), bson.M{})
	if err != nil {
		util.Lg.Debugf("find error: %v", err)
	} else {
		units := make([]robotUnit, 0)
		err = cur.All(context.TODO(), &units)
		if err != nil {
			util.Lg.Debugf("2. error: %v", err)
		}
		//util.Lg.Debugf("%v", len(units))
		nsid := RandIntRange(0, len(units) - 1)
		n    := RandIntRange(1, 3)
		if nsid >= len(units) {
			util.Lg.Debugf("%v, %v", nsid, n)
		}
	}
}

func main() {
	GetUnit()
}
