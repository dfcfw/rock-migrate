package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Action interface {
	Repository[model.Action]
}

func NewAction(db *mongo.Database) Action {
	repo := newBaseRepository[model.Action](db, "action")
	return &actionRepo{
		Repository: repo,
	}
}

type actionRepo struct {
	Repository[model.Action]
}
