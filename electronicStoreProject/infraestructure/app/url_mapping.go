package app

import (
	"../controllers"
)

func mapUrls(handler controllers.RedirectProductoHandler) {
	router.POST("/productos", handler.Create)

	router.GET("/productos/:codigo", handler.Get)

	router.GET("/productos", handler.FindAll)

	router.PUT("/productos/:codigo", handler.Update)

	router.DELETE("/productos/:codigo", handler.Delete)

}