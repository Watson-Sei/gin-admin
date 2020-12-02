package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var DB *gorm.DB

type DBConfig struct {
	Host 		string
	User 		string
	DBName		string
	Password 	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:		"tcp(db)",
		User:		os.Getenv("MYSQL_USER"),
		Password: 	os.Getenv("MYSQL_PASSWORD"),
		DBName: 	os.Getenv("MYSQL_DATABASE"),
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", DBUrl(BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	return db
}