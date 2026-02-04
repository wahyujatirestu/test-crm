package routes

import (
	"github.com/gin-gonic/gin"
	"test_crm/controllers"
	"test_crm/middleware"
)

func ContactRoutes(
	r *gin.RouterGroup,
	authMw middleware.AuthMiddleware,
	cc *controller.ContactController,
) {
	contacts := r.Group("/memberships/:id/contacts")
	contacts.Use(authMw.RequireToken())

	contacts.POST("", cc.Create)
	contacts.GET("", cc.GetByMembership)

	r.PUT("/contacts/:id", authMw.RequireToken(), cc.Update)
	r.DELETE("/contacts/:id", authMw.RequireToken(), cc.Delete)
}

