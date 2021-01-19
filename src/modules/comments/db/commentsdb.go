package comments

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentsDBInterface ..
type CommentsDBInterface interface {
	GetCommentsByIdea(query string) []MimicUser
	Insert(obj interface{}) error
}

// NewCommentsDB ..
func NewCommentsDB() CommentsDBInterface {
	return &CommentsDB{connection: "connection established"}
}

// CommentsDB ..
type CommentsDB struct {
	connection string
}


type Status string

const (
    Approved Status = "A",
    Spam = "X",
    Reported = "R",
    Flagged = "F"
)
// MimicComment Just to minic Comments collection
type MimicComment struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       primitive.ObjectID `bson:"userId,omitempty"`
	UserReported []string           `bson:"userReported,omitempty"`
	Uuid         string             `bson:"uuid,omitempty"`
	Comment      string             `bson:"comment,omitempty"`
	IdeaId       string             `bson:"ideaId,omitempty"`
	IsPublished  bool               `bson:"isPublished,omitempty"`
	IsSanitized  bool               `bson:"isSanitized,omitempty"`
	iIActive     bool               `bson:"isActive,omitempty"`
	Status       Status             `bson:"status,omitempty"`
	ParentId     string             `bson:"parentId,omitempty"`
	CreatedAt 	 time.Time 			`bson:"created_at"`
}


// Get ..
func (m *CommentsDB) GetCommentsByIdea(query string) []MimicComment {

	return []MimicSaSoftAction{{ID: "1", UserId: "Shepard"}, {ID: "2", UserId: "Miranda"}}
}

// Insert ..
func (m *CommentsDB) Insert(obj interface{}) error {
	return nil
}
