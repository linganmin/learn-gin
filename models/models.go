package models

import (
	"fmt"
	"ginhello/packages/setting"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn `json:"created_on"`
	ModifyOn `json:"modified_on"`
}

func init()  {
	var (
		err error
		dbType,dbName,username,password,host,port,tablePrefix string
	)

	sec,err := setting.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("DRIVER").String()
	dbName= sec.Key("USERNAME").String()
	password= sec.Key("PASSWORD").String()
	host= sec.Key("HOST").String()
	port= sec.Key("PORT").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()


	db,err := gorm.Open(dbType,fmt.Sprintf("&s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",username,password,host,port,dbName))

	if err != nil {

		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}