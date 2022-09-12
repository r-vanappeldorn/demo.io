package repositories

import (
	"context"
	"function-srv/pkg/database"
	"function-srv/pkg/models"
)

const collectionName = "functions"

type FunctionRepository interface {
	FindOne(context.Context, interface{}) (*models.Function, error)
	Create(context.Context, *models.Function) error
}

type functionRepository struct {
	db database.DatabaseHelper
}

func NewFunctionRepository(db database.DatabaseHelper) FunctionRepository {
	return &functionRepository{db}
}

func (f *functionRepository) FindOne(ctx context.Context, filter interface{}) (*models.Function, error) {
	function := &models.Function{}
	err := f.db.Collection(collectionName).FindOne(ctx, filter).Decode(function)
	if err != nil {
		return nil, err
	}

	return function, err
}

func (f *functionRepository) Create(ctx context.Context, fun *models.Function) error {
	_, err := f.db.Collection(collectionName).InsertOne(ctx, fun)
	return err
}