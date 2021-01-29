package comments

import (
	"go-boilerplate-api/config"
	"go-boilerplate-api/pkg/clients/db"

	"github.com/gin-gonic/gin"
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

func (service *Service) commentsDel(ctx *gin.Context) {

}

func (service *Service) commentPost(ctx *gin.Context) {

}

func (service *Service) commentsEdit(ctx *gin.Context) {

}

func (service *Service) reportSpam(ctx *gin.Context) {

}
