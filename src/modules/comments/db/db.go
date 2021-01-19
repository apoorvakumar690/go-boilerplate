package comments

import (
	"go-boilerplate-api/config"
)

// Instances ... contains the interface layer of the different dbs
type Instances struct {
	CommentsDB CommentsDBInterface
}

// NewInstance creates an instance of initialized DBInstances
func NewInstance(conf config.IConfig) (*Instances, error) {
	dbInstances := &Instances{}

	commentsDBInterface, err := initCommentsDB(conf)
	if err != nil {
		return nil, err
	}

	// Sets db instance
	dbInstances.CommentsDB = commentsDBInterface

	return dbInstances, nil
}

// Simulates the initialization of a db connection
func initCommentsDB(config config.IConfig) (CommentsDBInterface, error) {
	return NewCommentsDB(), nil
}
