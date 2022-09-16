package repositories

import (
	"account-service/pkg/database"
	"account-service/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const collectionName = "users"

type UserRepo interface {
	Create(context.Context, *models.User) (string, error)
	FindByUserNameAndPassword(context.Context, string, string) (*models.User, error)
}

type User struct {
	db database.DatabaseHelper
}

func NewUserRepo(db database.DatabaseHelper) UserRepo {
	return &User{db}
}

func (u *User) Create(ctx context.Context, usr *models.User) (string, error) {
	id, err := u.db.Collection(collectionName).InsertOne(ctx, usr)
	return id, err
}

func (u *User) FindByUserNameAndPassword(ctx context.Context, userName, hashedPassword string) (*models.User, error) {
	usr := &models.User{}
	err := u.db.Collection(collectionName).FindOne(ctx, bson.D{{Key: "userName", Value: userName}, {Key: "password", Value: hashedPassword}}).Decode(usr)
	if err != nil {
		return nil, err
	}

	return usr, err
}
