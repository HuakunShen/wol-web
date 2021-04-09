package models

type Mac struct {
	Id     string `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId string `json:"user_id" gorm:"foreignKey:CompanyRefer"`
	Name   string `json:"name"`
	Mac    string `json:"mac"`
}
