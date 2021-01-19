package likes

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

type SaSoftAction struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       primitive.ObjectID `bson:"userId,omitempty"`
	Type 		 string           `bson:"type,omitempty"`
	UserReqRef   string           `bson:"userReqRef,omitempty"`
	Status   	 int           `bson:"status,omitempty"`
	SecondaryRef string           `bson:"secondaryRef,omitempty"`
	CreatedAt 	 time.Time 			`bson:"created_at"`
}

func NewSaSoftAction() *SaSoftAction {
	return &SaSoftAction{
		uuid: shortuuid.New(),
		CreatedAt: time.Now(),
	}
}
