package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Nama string `gorm:"type:varchar(255)" json:"nama" validate:"required,min=5"`
}

func (Kategori) TableName() string {
	return "ref_kategori"
}
