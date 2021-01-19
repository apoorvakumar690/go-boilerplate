package Comments

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/lithammer/shortuuid/v3"	
)
type Status string

const (
    Approved Status = "A",
    Spam = "X",
    Reported = "R",
    Flagged = "F"
)

type Comment struct {
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

func NewComment() *Comment {
	return &Comment{
		uuid: shortuuid.New(),
		Status: "A"
		IsPublished: true,
		IsSanitized: true,
		isActive: true,
		CreatedAt: time.Now(),
	}
}
