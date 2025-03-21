package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Module interface {
	Repository[model.Module]
}

func NewModule(db *mongo.Database) Module {
	repo := newBaseRepository[model.Module](db, "module")
	return &moduleRepo{
		Repository: repo,
	}
}

type moduleRepo struct {
	Repository[model.Module]
}
