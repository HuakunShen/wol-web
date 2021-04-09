package models

type Mac struct {
	Id     uint `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId uint `json:"user_id" gorm:"foreignKey:CompanyRefer"`
	Name   string `json:"name"`
	Mac    string `json:"mac"`
}
