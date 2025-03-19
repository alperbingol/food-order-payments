package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents a food order
type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ItemName string             `bson:"item_name" json:"item_name"`
	Price    float64            `bson:"price" json:"price"`
	Status   string             `bson:"status" json:"status"`
}
