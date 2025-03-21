package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlackJA3 struct {
	ID        primitive.ObjectID `json:"id"         bson:"_id,omitempty"`
	JA3       string             `json:"ja3"        bson:"ja3"`
	Labels    []string           `json:"labels"     bson:"labels"`
	Clusters  []string           `json:"clusters"   bson:"clusters"`
	CreatedBy string             `json:"created_by" bson:"created_by,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	ExpiredAt time.Time          `json:"expired_at" bson:"expired_at,omitempty"`
}
