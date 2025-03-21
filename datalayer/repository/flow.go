package repository

import (
	"context"

	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Flow interface {
	Repository[model.Flow]
}

func NewFlow(db *mongo.Database) Flow {
	repo := newBaseRepository[model.Flow](db, "flow")
	return &flowRepo{
		Repository: repo,
	}
}

type flowRepo struct {
	Repository[model.Flow]
}

func (repo *flowRepo) CreateIndex(ctx context.Context) error {
	idx := mongo.IndexModel{Keys: bson.D{{Key: "time", Value: -1}}}
	_, err := repo.Indexes().CreateOne(ctx, idx)

	return err
}
