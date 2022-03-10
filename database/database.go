package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Mongo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (m *Mongo) Open() {
	srv := os.Getenv("MURI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(srv))
	if err != nil {
		panic(err)
	}
	m.Client = client
}

func (m *Mongo) Close() {
	if err := m.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func (m *Mongo) SetCollection(collection string) {
	m.Collection = m.Client.Database("shop").Collection(collection)
}

//findOne
func (m *Mongo) FindOne(b *bson.D) bson.M {
	var result bson.M
	err := m.Collection.FindOne(context.TODO(), b).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return bson.M{"message": "Not found document"}
	}
	if err != nil {
		panic(err)
	}
	return result
}
