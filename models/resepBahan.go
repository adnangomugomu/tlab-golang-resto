package models

import "gorm.io/gorm"

type ResepBahan struct {
	gorm.Model
	ResepID int    `gorm:"type:int(11);" json:"resep_id" validate:"required"`
	BahanID int    `gorm:"type:int(11);" json:"bahan_id" validate:"required"`
	Takaran string `gorm:"type:varchar(255)" json:"takaran" validate:"required"`
	Resep   Resep  `gorm:"foreignKey:ResepID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"resep" validate:"-"`
	Bahan   Bahan  `gorm:"foreignKey:BahanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"bahan" validate:"-"`
}

func (ResepBahan) TableName() string {
	return "resep_bahan"
}
