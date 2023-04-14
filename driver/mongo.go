package driver

import (
	"context"
	"fmt"

	"github.com/pcpratheesh/dbmux/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	db *mongo.Client
}

func mongoConn() *mongoClient {
	return &mongoClient{}
}

// initiate connection with mongo
func (mc *mongoClient) Connect(param entity.Options) error {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%v",
		param.User, param.Password, param.Host, param.Port)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mc.db = client

	return nil

}

func (ms *mongoClient) Close() error {
	return nil
}

// Pool
func (ms *mongoClient) Pool() interface{} {
	return ms.db
}
