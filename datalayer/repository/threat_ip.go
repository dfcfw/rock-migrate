package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
