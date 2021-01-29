package likes

import (
	"go-boilerplate-api/shared"
	likes "go-boilerplate-api/src/modules/likes/controller"

	"github.com/gin-gonic/gin"
)

// NewUserRoute Creates and initializes user routes
func NewUserRoute(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := likes.NewLikeService(deps.Config, deps.Database)
	userAPI := router.Group("/api/sa")
	{
		userAPI.POST("/like", service.likePost)
		userAPI.GET("/getLike/:uuid", service.getLikes)
	}
}
