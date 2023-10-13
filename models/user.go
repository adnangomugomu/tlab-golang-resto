package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string `gorm:"type:varchar(255)" json:"nama" validate:"required,min=5"`
	Username     string `gorm:"type:varchar(255);unique;not null" json:"username" validate:"required,min=5"`
	Password     string `gorm:"type:varchar(255)" json:"password" validate:"required,min=5"`
	FotoProfil   string `gorm:"type:varchar(255)" json:"foto_profil" validate:"-"`
	RefreshToken string `gorm:"type:text;column:refresh_token" json:"refresh_token" validate:"-"`
}

func (User) TableName() string {
	return "users"
}
