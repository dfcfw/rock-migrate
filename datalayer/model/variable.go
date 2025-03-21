package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Variable struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Label    string        `bson:"label,omitempty" json:"label"`
	Format   string        `bson:"format,omitempty" json:"format"`
	Name     string        `bson:"name,omitempty" json:"name"`
	Who      string        `bson:"who,omitempty" json:"who"`
	CreateAt time.Time     `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time     `bson:"update_at,omitempty" json:"update_at"`
}
