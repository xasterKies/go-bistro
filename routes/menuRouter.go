package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xasterKies/go-bistro/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controller.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menu", controller.CreateMenu())
	incomingRoutes.PATCH("/menu/menu_id", controller.UpdateMenu())
}