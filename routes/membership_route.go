package routes

import (
	"github.com/gin-gonic/gin"
	"test_crm/controllers"
	"test_crm/middleware"
)

func MembershipRoutes(
	r *gin.RouterGroup,
	authMw middleware.AuthMiddleware,
	mc *controller.MembershipController,
) {
	membership := r.Group("/memberships")
	membership.Use(authMw.RequireToken())

	membership.POST("", mc.Create)
	membership.GET("", mc.GetAll)

	membership.GET("/detail/:id", mc.GetByID)

	membership.PUT("/:id", mc.Update)
	membership.DELETE("/:id", mc.Delete)

	membership.GET("/with-contacts", mc.GetActiveWithContact)
}

