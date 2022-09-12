package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Function struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
	Salary float32 `json:"salary" bson:"salary"`
}