package comments

import (
	"go-boilerplate-api/shared"
	comments "go-boilerplate-api/src/modules/comments/controller"

	"github.com/gin-gonic/gin"
)

// NewUserRoute Creates and initializes user routes
func NewUserRoute(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	service := comments.NewCommentsService(deps.Config, deps.Database)
	userAPI := router.Group("/api/sa")
	{
		userAPI.POST("/comment", service.insert)
		userAPI.DELETE("/comment/delete", service.getAll)
		userAPI.PUT("/comment/edit", service.insert)
		userAPI.POST("/comment/report-spam", service.getAll)
	}
}
