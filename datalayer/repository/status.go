package repository

import (
	"context"

	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Status interface {
	Repository[model.Status]
}

func NewStatus(db *mongo.Database) Status {
	repo := newBaseRepository[model.Status](db, "status")
	return &statusRepo{
		Repository: repo,
	}
}

type statusRepo struct {
	Repository[model.Status]
}

func (repo *statusRepo) CreateIndex(ctx context.Context) error {
	idx := mongo.IndexModel{Keys: bson.D{{Key: "time", Value: -1}}}
	_, err := repo.Indexes().CreateOne(ctx, idx)

	return err
}
