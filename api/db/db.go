package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Hendryboyz/web3-api-go/configs"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	config := configs.GetConfig()
	connectString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetInt32("database.port"),
		config.GetString("database.name"),
	)
	db, err = gorm.Open(mysql.Open(connectString), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect fail")
	} else {
		log.Print("DB connected")
	}
}

func GetDBInstance() *gorm.DB {
	return db
}
