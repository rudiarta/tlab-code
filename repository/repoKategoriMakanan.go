package repository

import (
	"github.com/rudiarta/tlab-code/database"
	"github.com/rudiarta/tlab-code/model"
)

func AddKategoriMakanan(data *model.KategoriMakanan) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func CountByNameKategoriMakanan(name string) (int, error) {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	var count int
	data := model.KategoriMakanan{}
	count = 0
	if err := db.Model(&data).Where("name = ?", name).Count(&count).Error; err != nil {
		return 0, err
	}

	if err := db.Where("name = ?", name).First(&data).Error; err != nil {
		return 0, err
	}

	return data.ID, nil
}

func UpdateKategoriMakanan(id int, name string) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Model(model.KategoriMakanan{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}

	return nil
}

func ListKategoriMakanan(data *[]model.KategoriMakanan) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Find(&data).Error; err != nil {
		return err
	}

	return nil
}
