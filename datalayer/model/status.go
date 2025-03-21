package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Status struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Time5   int64         `bson:"time5,omitempty"`
	Time30  int64         `bson:"time30,omitempty"`
	Time60  int64         `bson:"time60,omitempty"`
	Uuid    string        `bson:"uuid,omitempty"`
	Type    string        `bson:"type,omitempty"`
	Name    string        `bson:"name,omitempty"`
	Cluster string        `bson:"cluster,omitempty"`
	Private string        `bson:"private,omitempty"`
	Time    time.Time     `bson:"time,omitempty"`
	Status  StatusItem    `bson:"status,omitempty"`
}

type StatusItem struct {
	Total  int64 `bson:"total,omitempty"  json:"total"`
	Deny   int64 `bson:"deny,omitempty"   json:"deny"`
	S200   int64 `bson:"_200,omitempty"   json:"_200"`
	S30x   int64 `bson:"_30x,omitempty"   json:"_30x"`
	S40x   int64 `bson:"_40x,omitempty"   json:"_40x"`
	S50x   int64 `bson:"_50x,omitempty"   json:"_50x"`
	S100   int64 `bson:"_100,omitempty"   json:"_100"`
	S1000  int64 `bson:"_1000,omitempty"  json:"_1000"`
	S10000 int64 `bson:"_10000,omitempty" json:"_10000"`
	Sx0000 int64 `bson:"_x0000,omitempty" json:"_x0000"`
}
