package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Security 安全插件
type Security struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"       json:"id"`
	Label     string             `bson:"label,omitempty"     json:"label"`
	Status    string             `bson:"status,omitempty"    json:"status"`
	Weight    int                `bson:"weight"              json:"weight"`
	Action    string             `bson:"action,omitempty"    json:"action"`
	Module    string             `bson:"module,omitempty"    json:"module"`
	Args      []interface{}      `bson:"args,omitempty"      json:"args"`
	App       []string           `bson:"app,omitempty"       json:"app"`
	Cluster   []string           `bson:"cluster,omitempty"   json:"cluster"`
	Condition []Condition        `bson:"condition"           json:"condition"`
	Extra     []Extra            `bson:"extra,omitempty"     json:"extra"`
	Who       string             `bson:"who,omitempty"       json:"who"`
	Pinyin    string             `bson:"pinyin,omitempty"    json:"-"`
	CreateAt  time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt  time.Time          `bson:"update_at,omitempty" json:"update_at"`
}

type Condition struct {
	Var    string   `bson:"var,omitempty"    json:"var"    binding:"required"`
	Method string   `bson:"method,omitempty" json:"method" binding:"oneof=regex pool equal prefix suffix contain cidr script 'not regex' 'not equal' 'not pool' 'not prefix' 'not suffix' 'not contain' 'not cidr' 'not script'"`
	Data   []string `bson:"data,omitempty"   json:"data"   binding:"gte=1,dive,required"`
}

type Extra struct {
	Name  string `json:"name"  bson:"name"  binding:"required"`
	Value bool   `json:"value" bson:"value"`
}
