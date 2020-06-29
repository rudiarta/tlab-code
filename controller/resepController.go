package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudiarta/tlab-code/model"
	"github.com/rudiarta/tlab-code/repository"
)

func AddResepController(c *gin.Context) {

	kategori := c.PostForm("kategori")
	resepName := c.PostForm("name")
	resepDescription := c.PostForm("description")

	kategoriID, err := repository.CountByNameKategoriMakanan(kategori)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Kategori not found.",
		})

		return
	}

	resep := model.Resep{
		KategoriID: kategoriID,
		Name:       resepName,
	}
	if err := repository.AddResep(&resep); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}
	bahan := model.BahanMakanan{
		ResepID:     resep.ID,
		Description: resepDescription,
	}
	if err := repository.AddBahanMakanan(&bahan); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "success.",
	})

	return
}

func DeleteResepController(c *gin.Context) {
	idResep, _ := strconv.Atoi(c.PostForm("id"))
	if err := repository.DeleteBahanMakananByResepID(idResep); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := repository.DeleteResep(idResep); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "success.",
	})

	return
}

func ListResepController(c *gin.Context) {
	data := []model.ListDetailResep{}
	if err := repository.ListResep(&data); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
	return
}

func UpdateResepController(c *gin.Context) {
	resepID, _ := strconv.Atoi(c.PostForm("id"))
	if err := repository.UpdateResep(resepID, c.PostForm("name"), c.PostForm("bahan")); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "success.",
	})
}

func GetResepController(c *gin.Context) {
	data := model.ListDetailResep{}
	resepID, _ := strconv.Atoi(c.Param("id"))
	if err := repository.GetResep(resepID, &data); err != nil {
		c.JSON(422, gin.H{
			"mesage": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "succes.",
		"data":    data,
	})
	return
}
