package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Upstream interface {
	Repository[model.Upstream]
}

func NewUpstream(db *mongo.Database) Upstream {
	repo := newBaseRepository[model.Upstream](db, "upstream")
	return &upstreamRepo{
		Repository: repo,
	}
}

type upstreamRepo struct {
	Repository[model.Upstream]
}
