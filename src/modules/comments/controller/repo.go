package comments

import (
	"go-boilerplate-api/config"
	likes "go-boilerplate-api/src/modules/likes/db"
)

// CommentRepoInterface ...
type CommentRepoInterface interface {
	Get(query string) ([]*User, error)
	GetOne(id string) (*User, error)
	GetAll() ([]*User, error)
	Insert(u User) error
}

// NewCommentRepo Create's an instance of a Comments Repository
func NewCommentRepo(conf config.IConfig, dbInstances *likes.Instances) CommentRepoInterface {
	return &CommentRepo{config: conf, db: dbInstances.CommentsDB}
}

// CommentRepo Contains methods to action on the Comment Repository
type CommentRepo struct {
	config config.IConfig
	db     likes.CommentsDBInterface
}

// GetCommentsByIdea Gets users using a query
func (ur *CommentRepo) GetCommentsByIdea(query string) ([]*Comments, error) {
	u := ur.db.GetCommentsByIdea(query)
	likes := bindToComments(u)
	return likes, nil
}

// Insert Inserts a User
func (ur *CommentRepo) Insert(u User) error {
	ur.db.Insert(u)
	return nil
}

func bindToComments(u []comments.MimicComment) []*Comments {
	Comment := []*Comments{}
	for i := 0; i < len(u); i++ {
		comment = append(user, bindToComment(u[i]))
	}
	return comment
}

func bindToComment(u comments.MimicComment) *Comments {
	comment := &Comments{
		ID: u.Id,          
		UserId: u.UserId,       
		UserReported: u.UserReported,  
		Uuid: u.Uuid,         
		Comment: u.Comments,      
		IdeaId: u.IdeaId,       
		IsPublished: u.IsPublished,  
		IsSanitized: u.IsSanitized  
		iIActive: u.iIActive     
		Status: u.Status       
		ParentId: u.ParentId     
		CreatedAt: u.CreatedAt
	}
	return comment
}
