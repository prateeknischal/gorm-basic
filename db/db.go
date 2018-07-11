package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:password@(:3306)/gorm_basic?charset=utf8&parseTime=True")
	if err != nil {
		panic(fmt.Sprintf("Failed to open a connection; err: %s", err.Error()))
	}
}
