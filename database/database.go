package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "root:maulapor@(localhost)/maulapor?charset=utf8&parseTime=True&loc=Local")
	return
}
