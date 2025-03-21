package repository

import (
	"context"

	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ThreatIP interface {
	Repository[model.ThreatIP]
}

func NewThreatIP(db *mongo.Database) ThreatIP {
	repo := newBaseRepository[model.ThreatIP](db, "threat_ip")
	return &threatIPRepo{
		Repository: repo,
	}
}

type threatIPRepo struct {
	Repository[model.ThreatIP]
}

func (repo *threatIPRepo) CreateIndex(ctx context.Context) error {
	idx := mongo.IndexModel{Keys: bson.D{{Key: "last_at", Value: -1}}}
	_, err := repo.Indexes().CreateOne(ctx, idx)

	return err
}
