package models

type Mac struct {
	Id     uint `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId uint `json:"user_id" gorm:"foreignKey:CompanyRefer"`
	Name   string `json:"name" gorm:"not null"`
	Mac    string `json:"mac" gorm:"not null"`
	Ip string `json:"ip" gorm:"not null"`
	Port uint32 `json:"port" gorm:"not null"`
}
