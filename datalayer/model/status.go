package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Time5   int64              `bson:"time5,omitempty"`
	Time30  int64              `bson:"time30,omitempty"`
	Time60  int64              `bson:"time60,omitempty"`
	Uuid    string             `bson:"uuid,omitempty"`
	Type    string             `bson:"type,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Cluster string             `bson:"cluster,omitempty"`
	Private string             `bson:"private,omitempty"`
	Time    time.Time          `bson:"time"`
	Status  StatusItem         `bson:"status"`
}

type StatusItem struct {
	Total  int64 `bson:"total"  json:"total"`
	Deny   int64 `bson:"deny"   json:"deny"`
	S200   int64 `bson:"_200"   json:"_200"`
	S30x   int64 `bson:"_30x"   json:"_30x"`
	S40x   int64 `bson:"_40x"   json:"_40x"`
	S50x   int64 `bson:"_50x"   json:"_50x"`
	S100   int64 `bson:"_100"   json:"_100"`
	S1000  int64 `bson:"_1000"  json:"_1000"`
	S10000 int64 `bson:"_10000" json:"_10000"`
	Sx0000 int64 `bson:"_x0000" json:"_x0000"`
}
