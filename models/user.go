package models

import (
	"github.com/Watson-Sei/gin-admin/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 			uint
	Username	string
	Password 	string
}

func CreateUser(username string, password string) (err error) {
	db := config.DBConnect()
	defer db.Close()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := db.Create(&User{Username: username, Password: string(hash)}).Error; err != nil {
		return err
	}
	return nil
}