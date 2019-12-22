package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type App struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App
	Org
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	GameId   string `json:"game_id"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(192.168.1.146:3306)/game?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("gorm open error: %v\n", err)
	}
	defer db.Close()

	ui := &UserInfo{
		Id:       0,
		UserId:   112233,
		UserName: "Jjjjww",
		GameId:   "600101",
	}

	//if err := db.Exec("insert into user_info (id, user_id, user_name, game_id) values (?, ?, ?, ?)", ui.Id, ui.UserId, ui.UserName, ui.GameId).Error; err != nil {
	//	log.Fatal(err)
	//}
	//ui2 := &UserInfo{}
	//n := db.Table("user_info").Where("user_id = ?", ui.UserId)
	n := db.Table("user_info").Find("user_id = ?", ui.UserId).RowsAffected
	//n := db.Raw("select * from user_info where user_id = ?", ui.UserId).Error
	log.Println("RowsAffected: ", n)
	//db.Create(ui)

	//log.Printf("result: %v\n", db.NewRecord(ui))

	//data := []byte(`
	//    {
	//        "id": "k34rAT4",
	//        "name": "My Awesome Org"
	//    }
	//`)
	//
	//var b AppWithOrg
	//
	//json.Unmarshal(data, &b)
	//fmt.Printf("1: %#v\n", b)
	//
	//a := AppWithOrg{
	//	App: App{
	//		Id: "k34rAT4",
	//	},
	//	Org: Org{
	//		Name: "My Awesome Org",
	//	},
	//}
	//data, _ = json.Marshal(a)
	//fmt.Println("2:", string(data))
}
