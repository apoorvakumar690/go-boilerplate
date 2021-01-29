package likes

import (
	"net/http"
	"strconv"

	"go-boilerplate-api/config"
	"go-boilerplate-api/pkg/clients/db"

	"github.com/gin-gonic/gin"
)

// Service contains the methods required to perfom operation's on users
type Service struct {
	like LikeRepoInterface
}

// NewLikeService ...
func NewLikeService(conf config.IConfig, db *db.Instances) *Service {
	likeRepo := NewLikeRepo(conf, db)
	return &Service{like: likeService}
}

func (service *Service) getLikes(ctx *gin.Context) {

	ideaUUID := ctx.Param("uuid")
	perPage := strconv.Atoi(ctx.DefaultQuery("perPage", "50"))
	page := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	likes, err := service.like.GetLikesByIdea()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

}

type likeContentSchema struct {
	LikeType         string `form:"likeType" json:"likeType" xml:"likeType"  binding:"required"`
	LikeDislikeValue string `form:"likeDislikeValue" json:"likeDislikeValue" xml:"likeDislikeValue" binding:"required"`
	UUID             string `form:"uuid" json:"uuid" xml:"uuid" binding:"uuid"`
}

func (service *Service) likePost(ctx *gin.Context) {

	var reqBody likeContentSchema
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, "Request is improper")
		return
	}
}
