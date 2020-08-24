package models

type User struct {
	Id       int    `gorm:"primary_key"`
	Username string `gorm:"size:255"`
	Count    int
}

type FlagsAuth struct {
	ConsumerKey    string `header:“ConsumerKey”`
	ConsumerSecret string `header:“ConsumerSecret”`
}
