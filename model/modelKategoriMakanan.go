package model

type BaseKategoriMakanan struct {
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;"`
}

type KategoriMakanan struct {
	BaseKategoriMakanan
	Name string `gorm:"column:name"`
}

func (KategoriMakanan) TableName() string {
	return "resep_kategori"
}
