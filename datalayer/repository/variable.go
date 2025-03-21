package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Variable interface {
	Repository[model.Variable]
}

func NewVariable(db *mongo.Database) Variable {
	repo := newBaseRepository[model.Variable](db, "variable")
	return &variableRepo{
		Repository: repo,
	}
}

type variableRepo struct {
	Repository[model.Variable]
}
