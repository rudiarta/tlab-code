package repository

import (
	"github.com/rudiarta/tlab-code/database"
	"github.com/rudiarta/tlab-code/model"
)

func AddResep(resepData *model.Resep) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Create(resepData).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResep(id int) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if err := db.Where("id = ?", id).Delete(model.Resep{}).Error; err != nil {
		return err
	}

	return nil
}

func ListResep(data *[]model.ListDetailResep) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	result := db.Table("resep")
	result = result.Select("resep.id as id, resep.name as name, resep_kategori.name as kategori_name, bahan_makanan.description as bahan")
	result = result.Joins("INNER JOIN bahan_makanan ON bahan_makanan.resep_id = resep.id INNER JOIN resep_kategori ON resep.kategori_id = resep_kategori.id")
	result.Scan(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateResep(resepID int, name string, bahan string) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	if name != "" {
		if err := db.Model(model.Resep{}).Where("id = ?", resepID).Update("name", name).Error; err != nil {
			return err
		}
	}

	if err := db.Model(model.BahanMakanan{}).Where("resep_id = ?", resepID).Update("description", bahan).Error; err != nil {
		return err
	}

	return nil
}

func GetResep(resepID int, data *model.ListDetailResep) error {
	db, err := database.InitDatabase()
	defer db.Close()
	if err != nil {
		return err
	}

	result := db.Table("resep")
	result = result.Select("resep.id as id, resep.name as name, resep_kategori.name as kategori_name, bahan_makanan.description as bahan")
	result = result.Where("resep.id = ?", resepID)
	result = result.Joins("INNER JOIN bahan_makanan ON bahan_makanan.resep_id = resep.id INNER JOIN resep_kategori ON resep.kategori_id = resep_kategori.id")
	result.Scan(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
