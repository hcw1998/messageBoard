package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "test"
	PASSWORD = "464"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "message_board"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
	})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
