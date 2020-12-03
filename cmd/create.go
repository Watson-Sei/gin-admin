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

var dsn = fmt.Sprintf("%s:%s@%s/%s", os.Getenv("MYSQL_USER"),os.Getenv("MYSQL_PASSWORD"),os.Getenv("HOST"),os.Getenv("MYSQL_DATABASE"))

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
				if err := CreateUser(name, password); err != nil {
					fmt.Println("Good Job")
				} else {
					fmt.Println(err)
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

func CreateUser(username string, password string) (err error) {
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println("Database Close")
		panic(err)
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := db.Create(&User{Username: username, Password: string(hash)}).Error; err != nil {
		return err
	}
	fmt.Println("処理が成功しました。")
	return nil
}