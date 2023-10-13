package models

import "gorm.io/gorm"

type Resep struct {
	gorm.Model
	Nama       string   `gorm:"type:varchar(255)" json:"nama" validate:"required,min=5"`
	Keterangan string   `gorm:"type:text" json:"keterangan" validate:"required,min=5"`
	KategoriID int      `gorm:"type:int(11);" json:"kategori_id" validate:"required"`
	Kategori   Kategori `gorm:"foreignKey:KategoriID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"kategori" validate:"-"`
}

func (Resep) TableName() string {
	return "resep"
}
