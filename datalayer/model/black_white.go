package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlackWhite struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	Action    string             `bson:"action,omitempty"    json:"action"`
	Label     string             `bson:"label,omitempty"     json:"label"`
	Cluster   []string           `bson:"cluster,omitempty"   json:"cluster"`
	Condition []Condition        `bson:"condition"           json:"condition"`
	CreateAt  time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt  time.Time          `bson:"update_at,omitempty" json:"update_at"`
	Who       string             `bson:"who,omitempty"       json:"who"`
	Pinyin    string             `bson:"pinyin,omitempty"    json:"-"`
	Priority  int                `bson:"priority,omitempty"  json:"priority"`
	Enable    bool               `bson:"enable"              json:"enable"`
}
