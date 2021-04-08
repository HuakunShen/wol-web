package models

type Mac struct {
	Name string `json:"name" gorm:"primary_key,autoIncrement:false"`
	Mac  string `json:"mac"`
}

