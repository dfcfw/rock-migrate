package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Module struct {
	ID           bson.ObjectID       `bson:"_id,omitempty"       json:"id"`
	Args         []map[string]string `bson:"args"                json:"args"`
	Code         bson.Binary         `bson:"code,omitempty"      json:"-"`
	CreateAt     time.Time           `bson:"create_at,omitempty" json:"create_at"`
	Data         bson.Binary         `bson:"data,omitempty"      json:"-"`
	Hash         string              `bson:"hash,omitempty"      json:"hash"`
	Label        string              `bson:"label,omitempty"     json:"label"`
	UpdateAt     time.Time           `bson:"update_at,omitempty" json:"update_at"`
	Name         string              `bson:"name,omitempty"      json:"name"`
	Phase        []string            `bson:"phase,omitempty"     json:"phase"`
	Cluster      []string            `bson:"cluster,omitempty"   json:"cluster"`
	Type         string              `bson:"type,omitempty"      json:"type"`
	Who          string              `bson:"user,omitempty"      json:"who"`
	Remark       string              `bson:"remark"              json:"remark"`
	DisplayStyle string              `bson:"display_style"       json:"display_style"`
	Document     string              `bson:"document"            json:"document"` // 文档
}
