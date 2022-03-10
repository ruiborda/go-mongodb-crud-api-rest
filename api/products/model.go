package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID          primitive.ObjectID   `bson:"_id" json:"id"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	Price       primitive.Decimal128 `bson:"price" json:"price"`
	CreatedAt   primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}
