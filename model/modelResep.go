package model

type BaseResep struct {
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
}

type Resep struct {
	BaseResep
	KategoriID int    `gorm:"column:kategori_id;"`
	Name       string `gorm:"column:name;"`
}

func (Resep) TableName() string {
	return "resep"
}

type ListDetailResep struct {
	ID           int    `gorm:"column:id;" json:"id"`
	Name         string `gorm:"column:name;" json:"name"`
	KategoriName string `gorm:"column:kategori_name;" json:"kategori_name"`
	Bahan        string `gorm:"column:bahan;" json:"bahan"`
}
