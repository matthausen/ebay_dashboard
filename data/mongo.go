package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "mongodb+srv://matteo:<password>@ebay-dashboard-4twv3.gcp.mongodb.net/test?retryWrites=true&w=majority", nil)
