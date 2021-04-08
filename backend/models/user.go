package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}

type Mac struct {
	Name string `json:"name" gorm:"primary_key,autoIncrement:false"`
	Mac  string `json:"mac"`
}
