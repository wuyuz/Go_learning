package common

import (
	"context"
	"data_worker/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	DATABASE_URL string
}

var DBConfig *MongoDb

func loadDbConfig() {
	ViperConfig.SetDefault("DATABASE_URL", "")
	ViperConfig.Unmarshal(&DBConfig)
}

// GetMongoClient returns *mongo.Client
func GetMongoClient() (*mongo.Client, error) {
	loadDbConfig()
	return getMongoClientByURI(DBConfig.DATABASE_URL)
}

func getMongoClientByURI(uri string) (*mongo.Client, error) {
	var (
		err    error
		client *mongo.Client
		opts   *options.ClientOptions
	)
	if app.Client != nil {
		return app.Client, nil
	}

	opts = options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(10)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		return client, err
	}
	client.Ping(context.Background(), nil)
	return client, err
}
