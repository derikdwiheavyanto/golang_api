package app

import (
	"golang_restful_api/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router, controller controller.CategoryController) {
	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.PUT("/api/categories", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)
}
