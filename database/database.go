package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (m *Mongo) Open() {
	srv := os.Getenv("MURI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(srv))
	if err != nil {
		panic(err)
	}
	m.Client = client
	m.Database = m.Client.Database("shop")
}

func (m *Mongo) Close() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
