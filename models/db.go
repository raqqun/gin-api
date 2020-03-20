package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// InitDB hundler for db connection and db migration
func InitDB() {
	var err error
	dbHost := os.Getenv("MYSQL_HOST")
	dbURI := fmt.Sprintf("root:root@(%s)/ginapi?charset=utf8&parseTime=True&loc=Local", dbHost)
	db, err = gorm.Open("mysql", dbURI)
	if err != nil {
		log.Printf("Cannot connect to %s database\n", "mysql")
		log.Panicln("Error", err)
	}

	gormDebug, err := strconv.ParseBool(os.Getenv("GORM_DEBUG"))
	if err != nil {
		log.Println(
			"GORM_DEBUG expects boolean 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.",
			"Used false by default.")
		gormDebug = true
	}

	db.LogMode(gormDebug)

	db.Debug().AutoMigrate(
		&Users{},
	)
}
