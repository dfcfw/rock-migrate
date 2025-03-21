package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type White interface {
	Repository[model.BlackWhite]
}

func NewWhite(db *mongo.Database) White {
	repo := newBaseRepository[model.BlackWhite](db, "white")
	return &whiteRepo{
		Repository: repo,
	}
}

type whiteRepo struct {
	Repository[model.BlackWhite]
}
