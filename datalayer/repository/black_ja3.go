package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BlackJA3 interface {
	Repository[model.BlackJA3]
}

func NewBlackJA3(db *mongo.Database) BlackJA3 {
	repo := newBaseRepository[model.BlackJA3](db, "black_ja3")
	return &blackJA3Repo{
		Repository: repo,
	}
}

type blackJA3Repo struct {
	Repository[model.BlackJA3]
}
