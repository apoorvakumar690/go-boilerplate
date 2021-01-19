package comments

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
	comment CommentRepoInterface
}

// NewCommentService ...
func NewCommentsService(conf config.IConfig, db *db.Instances) *Service {
	commentRepo := NewCommentRepo(conf, db)
	return &Service{comment: commentService}
}

func (service *Service) getCommentsByIdea(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	ideaID := ctx.Param("uuid")
	comments, err := service.comment.GetCommentsByIdea(ideaID)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (service *Service) insert(ctx *gin.Context) {
	var err error
	defer utils.HandleError(ctx, &err)

	var comment comments.SaSoftAction
	if err = ctx.ShouldBindJSON(&comment); err != nil {
		err = errors.NewBadRequest("Could not bind request to model").SetCode("APIS.HTTP.USER.REQUEST_BIND_FAILD")
		return
	}

	err = service.comment.Insert(comment)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
