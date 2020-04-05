package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

var Client, err = mongo.Connect(context.Background(), "mongodb+srv://matteo:@eqG1t!s66Qc@ebay-dashboard-4twv3.gcp.mongodb.net/test?retryWrites=true&w=majority", nil)
