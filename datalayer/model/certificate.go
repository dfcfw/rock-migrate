package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certificate struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"          json:"id"`           // ID
	Name         string             `bson:"name,omitempty"         json:"name"`         // 名字
	Label        string             `bson:"label"                  json:"label"`        // 标签, 描述
	Pem          primitive.Binary   `bson:"pem,omitempty"          json:"-"`            // 证书
	Key          primitive.Binary   `bson:"key,omitempty"          json:"-"`            // 密钥
	PemHash      string             `bson:"phash,omitempty"        json:"-"`            // 证书Hash
	KeyHash      string             `bson:"khash"                  json:"-"`            // 密钥hash
	Expires      string             `bson:"expires,omitempty"      json:"expires"`      // 过期时间 (旧数据)
	Valid        string             `bson:"valid,omitempty"        json:"valid"`        // 有效期 (旧数据)
	DNSNames     []string           `bson:"dns_names,omitempty"    json:"dns_names"`    // 作用的域名
	NotBefore    time.Time          `bson:"not_before,omitempty"   json:"not_before"`   // 不早于, 即证书生效时间
	NotAfter     time.Time          `bson:"not_after,omitempty"    json:"not_after"`    // 不晚于, 即证书失效时间
	Country      []string           `bson:"country,omitempty"      json:"country"`      // 国家: 例如 CN
	Province     []string           `bson:"province,omitempty"     json:"province"`     // 省份: 例如 上海市
	Organization []string           `bson:"organization,omitempty" json:"organization"` // 组织: 例如 EastMoney Information Co.,Ltd
	CommonName   string             `bson:"common_name,omitempty"  json:"common_name"`  // 公用名: 例如 *.eastmoney.com
	CreateAt     time.Time          `bson:"create_at,omitempty"    json:"create_at"`    // 创建时间
	UpdateAt     time.Time          `bson:"update_at,omitempty"    json:"update_at"`    // 修改时间
}
