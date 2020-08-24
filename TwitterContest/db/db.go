package db

import (
	"TwitterContest/types"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB
var err error

var dbUserName = "admin"
var dbPassword = "9160266544135"
var dbName = "testDb"

var user models.User

func DatabaseSQL(count *map[string]int) models.User {

	DB.Debug().DropTableIfExists(&models.User{})
	//Drops table if already exists
	DB.Debug().AutoMigrate(&models.User{})
	//Auto create table based on

	fmt.Println(len(*count))

	// Inserting data into database
	var max = 0
	for key, value := range *count {
		var temp models.User
		temp.Username = key
		temp.Count = value
		//fmt.Println(user.Username, " ", user.Count)
		DB.Create(&temp)
		if max < value {
			max = value
		}
	}

	DB.Where("Count = ?", max).First(&user)
	//SELECT * FROM users WHERE count = max;
	fmt.Println("Hello")

	return user
}

func ConnectDb() {
	DB, err = gorm.Open("mysql", dbUserName+":"+dbPassword+"@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local")

	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

}
