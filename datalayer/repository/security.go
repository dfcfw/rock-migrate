package repository

import (
	"github.com/dfcfw/rock-migrate/datalayer/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Security interface {
	Repository[model.Security]
}

func NewSecurity(db *mongo.Database) Security {
	repo := newBaseRepository[model.Security](db, "security")
	return &securityRepo{
		Repository: repo,
	}
}

type securityRepo struct {
	Repository[model.Security]
}
