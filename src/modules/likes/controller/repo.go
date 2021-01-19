package likes

import (
	"go-boilerplate-api/config"
	likes "go-boilerplate-api/src/modules/likes/db"
)

// LikeRepoInterface ...
type LikeRepoInterface interface {
	Get(query string) ([]*User, error)
	GetOne(id string) (*User, error)
	GetAll() ([]*User, error)
	Insert(u User) error
}

// NewLikeRepo Create's an instance of a Likes Repository
func NewLikeRepo(conf config.IConfig, dbInstances *likes.Instances) LikeRepoInterface {
	return &LikeRepo{config: conf, db: dbInstances.LikesDB}
}

// LikeRepo Contains methods to action on the Like Repository
type LikeRepo struct {
	config config.IConfig
	db     likes.LikesDBInterface
}

// GetLikesByIdea Gets users using a query
func (ur *LikeRepo) GetLikesByIdea(query string) ([]*SaSoftAction, error) {
	u := ur.db.GetLikesByIdea(query)
	likes := bindToLikes(u)
	return likes, nil
}

// Insert Inserts a User
func (ur *LikeRepo) Insert(u User) error {
	ur.db.Insert(u)
	return nil
}

func bindToLikes(u []likes.MimicSaSoftAction) []*SaSoftAction {
	user := []*User{}
	for i := 0; i < len(u); i++ {
		user = append(user, &User{ID: u[i].ID, Name: u[i].Name})
	}
	return user
}
