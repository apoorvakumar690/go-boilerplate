package likes

import (
	"net/http"

	"go-boilerplate-api/apis/http/utils"
	"go-boilerplate-api/config"
	"go-boilerplate-api/pkg/clients/db"

	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-errors"
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

func (service *Service) getLikesByIdea(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	ideaID := ctx.Param("uuid")
	likes, err := service.like.GetLikesByIdea(ideaID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, likes)
}

func (service *Service) insert(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	var like likes.SaSoftAction
	if err = ctx.ShouldBindJSON(&like); err != nil {
		err = errors.NewBadRequest("Could not bind request to model").SetCode("APIS.HTTP.USER.REQUEST_BIND_FAILD")
		return
	}

	err = service.like.Insert(like)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
