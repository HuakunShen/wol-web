package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}

