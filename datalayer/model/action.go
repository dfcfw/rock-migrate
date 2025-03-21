package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Action struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	Label    string             `bson:"label,omitempty"     json:"label"`
	Name     string             `bson:"name,omitempty"      json:"name"`
	Who      string             `bson:"who,omitempty"       json:"who"`
	CreateAt time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time          `bson:"update_at,omitempty" json:"update_at"`
}
