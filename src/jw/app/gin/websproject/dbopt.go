package main

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

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

func (do *dbObj) CreateTable() (err error) {

	if err = do.Db.Create(&sbSubItem{}).Error; err != nil {
		lg.Errorf("create sbSubItem table failed: %v", err)
	}
	if err = do.Db.Create(&pageItem{}).Error; err != nil {
		lg.Errorf("create pageItem table failed: %v", err)
	}
	if err = do.Db.Create(&pageContent{}).Error; err != nil {
		lg.Errorf("create pageContent table failed: %v", err)
	}
	if err = do.Db.Create(&sidebarMainItem{}).Error; err != nil {
		lg.Errorf("create sidebarMainItem table failed: %v", err)
	}
	if err = do.Db.Create(&sidebar{}).Error; err != nil {
		lg.Errorf("create sidebar table failed: %v", err)
	}

	return
}

