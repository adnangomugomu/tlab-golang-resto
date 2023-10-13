package models

import "gorm.io/gorm"

type Bahan struct {
	gorm.Model
	Nama string `gorm:"type:varchar(255)" json:"nama" validate:"required,min=5"`
}

func (Bahan) TableName() string {
	return "ref_bahan"
}
