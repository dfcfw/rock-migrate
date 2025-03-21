package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Black interface {
	Repository[model.BlackWhite]
}

func NewBlack(db *mongo.Database) Black {
	repo := newBaseRepository[model.BlackWhite](db, "black")
	return &blackWhiteRepo{
		Repository: repo,
	}
}

type blackWhiteRepo struct {
	Repository[model.BlackWhite]
}
