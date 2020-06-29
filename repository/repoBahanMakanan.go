package repository

import (
	"github.com/rudiarta/tlab-code/database"
	"github.com/rudiarta/tlab-code/model"
)

func AddBahanMakanan(data *model.BahanMakanan) error {
	db, err := database.InitDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBahanMakananByResepID(id int) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Where("resep_id = ?", id).Delete(model.BahanMakanan{}).Error; err != nil {
		return err
	}

	return nil
}
