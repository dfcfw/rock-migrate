package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ThreatIP struct {
	ID     bson.ObjectID   `bson:"_id,omitempty"     json:"id,omitempty"`
	IP     string          `bson:"ip,omitempty"      json:"ip,omitempty"`
	LastAt time.Time       `bson:"last_at,omitempty" json:"last_at,omitempty"`
	Rules  []*ThreatIPRule `bson:"rules,omitempty"   json:"rules,omitempty"`
}

type ThreatIPRule struct {
	App  string `bson:"app,omitempty"  json:"app,omitempty"`
	From string `bson:"from,omitempty" json:"from,omitempty"`
	Rule string `bson:"rule,omitempty" json:"rule,omitempty"`
}
