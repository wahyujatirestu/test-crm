package routes

import (
	"github.com/gin-gonic/gin"
	"test_crm/controllers"
)

func AuthRoutes(r *gin.RouterGroup, ac *controller.AuthController) {
	auth := r.Group("/auth")

	auth.POST("/login", ac.Login)
}
