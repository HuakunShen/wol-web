package models

type User struct {
	Id       uint   `json:"id"`
	Username     string `json:"username" gorm:"unique"`
	Password []byte `json:"-"`
}
