/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"os"
)

type User struct {
	ID			uint
	Username	string
	Password	string
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name, err := cmd.PersistentFlags().GetString("name"); err == nil {
			if password, err := cmd.PersistentFlags().GetString("password"); err == nil {
				err := CreateUser(name, password)
				if err != nil {
					fmt.Println("Good Job")
				} else {
					fmt.Println("Bad Err")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringP("name", "u", "", "create username")
	createCmd.PersistentFlags().StringP("password", "p", "", "create user password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Database Connect
var DB *gorm.DB

type DBConfig struct {
	Host 		string
	User		string
	DBName		string
	Password	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host: "tcp(db)",
		User: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
	return &dbConfig
}

func DBUul(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", DBUul(BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	return db
}

func CreateUser(username string, password string) (err error) {
	db := DBConnect()
	defer db.Close()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := db.Create(&User{Username: username, Password: string(hash)}).Error; err != nil {
		return err
	}
	return nil
}



