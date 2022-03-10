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

/*func (m *Mongo) FindOne(b *bson.D) (bson.M, error) {
	var result bson.M
	collection := m.Database.Collection("products")
	err := collection.FindOne(context.TODO(), b).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return bson.M{"message": "Not found document"}, nil
	}
	if err != nil {
		panic(err)
	}
	return result, err
}

func (m *Mongo) InsertOne(product products.Model) (*mongo.InsertOneResult, error) {
	collection := m.Database.Collection("products")
	result, err := collection.InsertOne(context.TODO(), product)
	return result, err
}
*/
