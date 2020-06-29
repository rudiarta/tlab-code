package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudiarta/tlab-code/model"
	"github.com/rudiarta/tlab-code/repository"
)

func AddKategoriController(c *gin.Context) {

	data := model.KategoriMakanan{
		Name: c.PostForm("name"),
	}

	if err := repository.AddKategoriMakanan(&data); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Success.",
	})
	return
}

func UpdateKategoriController(c *gin.Context) {
	kategoriID, _ := strconv.Atoi(c.PostForm("id"))

	if err := repository.UpdateKategoriMakanan(kategoriID, c.PostForm("name")); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})

	return
}

func ListKategoriController(c *gin.Context) {
	listDataKategori := []model.KategoriMakanan{}
	if err := repository.ListKategoriMakanan(&listDataKategori); err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    listDataKategori,
	})

	return
}
