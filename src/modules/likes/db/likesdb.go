package likes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LikesDBInterface ..
type LikesDBInterface interface {
	GetLikesByIdea(query string) []MimicUser
	Insert(obj interface{}) error
}

// NewLikesDB ..
func NewLikesDB() LikesDBInterface {
	return &LikesDB{connection: "connection established"}
}

// LikesDB ..
type LikesDB struct {
	connection string
}

// MimicSaSoftAction Just to minic Likes collection
type MimicSaSoftAction struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserId       primitive.ObjectID `bson:"userId,omitempty"`
	Type         string             `bson:"type,omitempty"`
	UserReqRef   string             `bson:"userReqRef,omitempty"`
	Status       int                `bson:"status,omitempty"`
	SecondaryRef string             `bson:"secondaryRef,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
}

// Get ..
func (m *LikesDB) GetLikesByIdea(query string) []MimicSaSoftAction {

	return []MimicSaSoftAction{{ID: "1", UserId: "Shepard"}, {ID: "2", UserId: "Miranda"}}
}

// Insert ..
func (m *LikesDB) Insert(obj interface{}) error {
	return nil
}
