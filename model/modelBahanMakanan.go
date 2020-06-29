package model

type BaseBahanMakanan struct {
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
}

type BahanMakanan struct {
	ResepID     int    `gorm:"column:resep_id;"`
	Description string `gorm:"column:description;"`
}

func (BahanMakanan) TableName() string {
	return "bahan_makanan"
}
