package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/tlab-code/controller"
	"github.com/rudiarta/tlab-code/middleware"
)

func InitRoutes(app *gin.Engine) {

	// General Middleware
	app.Use(middleware.CORSMiddleware)
	router := app

	resepRouter := router.Group("/resep")
	{
		resepRouter.POST("/add", controller.AddResepController)
		resepRouter.POST("/delete", controller.DeleteResepController)
		resepRouter.PUT("/update", controller.UpdateResepController)
		resepRouter.GET("/list", controller.ListResepController)
		resepRouter.GET("/get/:id", controller.GetResepController)
	}

	kategoriRouter := router.Group("/kategori")
	{
		kategoriRouter.POST("/add", controller.AddKategoriController)
		kategoriRouter.PUT("/update", controller.UpdateKategoriController)
		kategoriRouter.GET("/list", controller.ListKategoriController)
	}
}
