package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "Mongo DB URI here", nil)
