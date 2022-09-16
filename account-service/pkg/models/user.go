package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" binding:"required,email" bson:"email"`
	FullName string             `json:"fullName" binding:"required" bson:"fullName"`
	Username string             `json:"username" binding:"required" bson:"username"`
	Password string             `json:"password" binding:"required,min=8" bson:"password"`
}

func NewUser(email, fullName, userName, password string) *User {
	return &User{ID: primitive.NewObjectID(), Email: email, FullName: fullName, Username: userName, Password: password}
}
