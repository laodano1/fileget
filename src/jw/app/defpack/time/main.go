package main

import (
	"github.com/davyxu/golog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/go-sql-driver/mysql"
	"unsafe"
)

var sidebarData sidebar
var (
	lg = golog.New("my-backend")
	exeAbsPath string
	myDB *dbObj
)


type pageItem struct {
	//gorm.Model
	Name string `json:"name"`
	Href  string   `json:"href"`
}

type pageContent struct {
	PageObjs  []pageItem `json:"pageObjs"`
}

type sbSubItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type sidebarMainItem struct {
	Name     string        `json:"name"`
	SubItems []sbSubItem   `json:"subItems"`
	PageCnt  []pageContent `json:"pageCnt"`
}

type sidebar struct {
	List  []sidebarMainItem `json:"list"`
}

type dbObj struct {
	Db *gorm.DB

}

func main() {
	var err error
	var dbType string
	dbType = "mysql"
	//dbType = "sqlite3"
	//myDB, err = NewDBObj(dbType, "myweb.db")
	myDB, err = NewDBObj(dbType, "root:123456@tcp(10.0.0.32:3306)/tst")
	if err != nil {
		return
	}

	if dbType == "sqlite3" {
		myDB.CreateSQLite3Table()
	} else {
		myDB.CreateMysqlTable()
	}

}

func NewDBObj(dbType, dbHost string) (*dbObj, error) {
	db, err := gorm.Open(dbType, dbHost)
	if err != nil {
		lg.Errorf("open %v db failed: %v", dbType, err)
		return nil, err
	}
	return &dbObj{
		Db: db,
	}, nil
}

func (do *dbObj) CreateMysqlTable() (err error) {

	subIList := make([]sbSubItem, 0)
	subIList = append(subIList, sbSubItem{Name: "mp3", Href: "/mp3"})
	subIList = append(subIList, sbSubItem{Name: "mp4", Href: "/mp4"})
	subIList = append(subIList, sbSubItem{Name: "mkv", Href: "/mkv"})


	pgItems1 := make([]pageItem, 0)
	pgItems1 = append(pgItems1, pageItem{})
	pgcnt1 := pageContent{PageObjs: pgItems1}

	pgItems2 := make([]pageItem, 0)
	pgItems2 = append(pgItems2, pageItem{Name: "tesla", Href: "/video/tesla.mp4"})
	pgcnt2 := pageContent{PageObjs: pgItems2}

	pgItems3 := make([]pageItem, 0)
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgcnt3 := pageContent{PageObjs: pgItems3}

	pgcList := make([]pageContent, 0)
	pgcList = append(pgcList, pgcnt1)
	pgcList = append(pgcList, pgcnt2)
	pgcList = append(pgcList, pgcnt3)

	sbmi := sidebarMainItem{
	Name:     "Media",
	SubItems: subIList,
	PageCnt:  pgcList,
	}

	sbmiList := make([]sidebarMainItem, 0)
	sbmiList = append(sbmiList, sbmi)
	sbObj := &sidebar{
	List: sbmiList,
	}

	sidebarData = *sbObj

	var i int
	if err = do.Db.Table("INFORMATION_SCHEMA.TABLES").Where("TABLE_NAME = ?", "pageItem").Count(&i).Error; err != nil {
		lg.Errorf("get pageItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create pageItem", i)
		if i <= 0 {
			if err = do.Db.Table("pageItem").CreateTable(&pageItem{}).Error; err != nil {
				lg.Errorf("create pageItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("INFORMATION_SCHEMA.TABLES").Where("TABLE_NAME = ?", "pageContent").Count(&i).Error; err != nil {
		lg.Errorf("get pageContent count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create pageContent", i)
		if i <= 0 {
			if err = do.Db.Table("pageContent").CreateTable(&pageContent{}).Error; err != nil {
				lg.Errorf("create pageContent table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("INFORMATION_SCHEMA.TABLES").Where("TABLE_NAME = ?", "sbSubItem").Count(&i).Error; err != nil {
		lg.Errorf("get sbSubItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sbSubItem", i)
		if i <= 0 {
			if err = do.Db.Table("sbSubItem").CreateTable(&sbSubItem{}).Error; err != nil {
				lg.Errorf("create sbSubItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("INFORMATION_SCHEMA.TABLES").Where("TABLE_NAME = ?", "sidebarMainItem").Count(&i).Error; err != nil {
		lg.Errorf("get sidebarMainItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sidebarMainItem", i)
		if i <= 0 {
			if err = do.Db.Table("sidebarMainItem").CreateTable(&sidebarMainItem{}).Error; err != nil {
				lg.Errorf("create sidebarMainItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("INFORMATION_SCHEMA.TABLES").Where("TABLE_NAME = ?", "sidebar").Count(&i).Error; err != nil {
		lg.Errorf("get sidebar count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sidebar", i)
		if i <= 0 {
			if err = do.Db.Table("sidebar").CreateTable(&sidebar{}).Error; err != nil {
				lg.Errorf("create sidebar table failed: %v", err)
			}
		}
	}



	return
}



func (do *dbObj) CreateSQLite3Table() (err error) {

	subIList := make([]sbSubItem, 0)
	subIList = append(subIList, sbSubItem{Name: "mp3", Href: "/mp3"})
	subIList = append(subIList, sbSubItem{Name: "mp4", Href: "/mp4"})
	subIList = append(subIList, sbSubItem{Name: "mkv", Href: "/mkv"})


	pgItems1 := make([]pageItem, 0)
	pgItems1 = append(pgItems1, pageItem{})
	pgcnt1 := pageContent{PageObjs: pgItems1}

	pgItems2 := make([]pageItem, 0)
	pgItems2 = append(pgItems2, pageItem{Name: "tesla", Href: "/video/tesla.mp4"})
	pgcnt2 := pageContent{PageObjs: pgItems2}

	pgItems3 := make([]pageItem, 0)
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgcnt3 := pageContent{PageObjs: pgItems3}

	pgcList := make([]pageContent, 0)
	pgcList = append(pgcList, pgcnt1)
	pgcList = append(pgcList, pgcnt2)
	pgcList = append(pgcList, pgcnt3)

	sbmi := sidebarMainItem{
		Name:     "Media",
		SubItems: subIList,
		PageCnt:  pgcList,
	}

	sbmiList := make([]sidebarMainItem, 0)
	sbmiList = append(sbmiList, sbmi)
	sbObj := &sidebar{
		List: sbmiList,
	}

	sidebarData = *sbObj

	var i int
	if err = do.Db.Table("sqlite_master").Where("type = 'table' and name = ?", "pageItem").Count(&i).Error; err != nil {
		lg.Errorf("get pageItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create pageItem", i)
		if i <= 0 {
			if err = do.Db.Table("pageItem").CreateTable(&pageItem{}).Error; err != nil {
				lg.Errorf("create pageItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("sqlite_master").Where("type = 'table' and name = ?", "pageContent").Count(&i).Error; err != nil {
		lg.Errorf("get pageContent count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create pageContent", i)
		if i <= 0 {
			if err = do.Db.Table("pageContent").CreateTable(&pageContent{}).Error; err != nil {
				lg.Errorf("create pageContent table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("sqlite_master").Where("type = 'table' and name = ?", "sbSubItem").Count(&i).Error; err != nil {
		lg.Errorf("get sbSubItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sbSubItem", i)
		if i <= 0 {
			if err = do.Db.Table("sbSubItem").CreateTable(&sbSubItem{}).Error; err != nil {
				lg.Errorf("create sbSubItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("sqlite_master").Where("type = 'table' and name = ?", "sidebarMainItem").Count(&i).Error; err != nil {
		lg.Errorf("get sidebarMainItem count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sidebarMainItem", i)
		if i <= 0 {
			if err = do.Db.Table("sidebarMainItem").CreateTable(&sidebarMainItem{}).Error; err != nil {
				lg.Errorf("create sidebarMainItem table failed: %v", err)
			}
		}
	}

	if err = do.Db.Table("sqlite_master").Where("type = 'table' and name = ?", "sidebar").Count(&i).Error; err != nil {
		lg.Errorf("get sidebar count failed: %v", err)
	} else {
		lg.Debugf("i: %v, create sidebar", i)
		if i <= 0 {
			if err = do.Db.Table("sidebar").CreateTable(&sidebar{}).Error; err != nil {
				lg.Errorf("create sidebar table failed: %v", err)
			}
		}
	}

	return
}