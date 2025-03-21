package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BlackIP interface {
	Repository[model.BlackIP]
}

func NewBlackIP(db *mongo.Database) BlackIP {
	repo := newBaseRepository[model.BlackIP](db, "black_ip")
	return &blackIPRepo{
		Repository: repo,
	}
}

type blackIPRepo struct {
	Repository[model.BlackIP]
}
