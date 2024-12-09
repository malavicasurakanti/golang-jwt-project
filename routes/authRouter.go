package routes

import (
	controller "golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())

}
