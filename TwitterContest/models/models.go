package models

type flagsAuth struct {
	ConsumerKey    string `header:“ConsumerKey”`
	ConsumerSecret string `header:“ConsumerSecret”`
}

type User struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:"size:255"`
	Count    int
}
