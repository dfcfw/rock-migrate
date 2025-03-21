package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Upstream struct {
	ID           bson.ObjectID `bson:"_id,omitempty"           json:"id"`
	Name         string        `bson:"name,omitempty"          json:"name"`
	Label        string        `bson:"label,omitempty"         json:"label"`
	Scheme       string        `bson:"scheme,omitempty"        json:"scheme"`
	Method       string        `bson:"method,omitempty"        json:"method"`
	Host         string        `bson:"host,omitempty"          json:"host"`
	Page         string        `bson:"page,omitempty"          json:"page"`
	Peers        []string      `bson:"peers,omitempty"         json:"peers"`
	BackoffPeers []string      `bson:"backoff_peers,omitempty" json:"backoff_peers"`
	Switch       string        `bson:"switch,omitempty"        json:"switch"`
	Timeout      int64         `bson:"timeout,omitempty"       json:"timeout"`
	Who          string        `bson:"who,omitempty"           json:"who"`
	CreateAt     time.Time     `bson:"create_at,omitempty"     json:"create_at"`
	UpdateAt     time.Time     `bson:"update_at,omitempty"     json:"update_at"`
}
