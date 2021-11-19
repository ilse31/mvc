package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nik      int    `gorm:"size:200" json:"nik"`
	Name     string `gorm:"size:150" json:"name" form:"name"`
	Alamat   string `gorm:"size:100" json:"alamat"`
	Email    string `gorm:"size:100" json:"email" form:"email"`
	Password string `gorm:"size:100" json:"password" form:"password"`
	Token    string `gorm:"size:100" json:"token"`
}

type Admin struct {
	gorm.Model
	CreditCard CreditCard `gorm:"polymorphic:IdKab;"`
	// use UserName as foreign key
}

type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
	IdKab    int `gorm:"index"`
}

type Kabs struct {
	gorm.Model
	Kabupaten string
	CrediCard CreditCard `gorm:"polymorphic:IdKab;"`
}
