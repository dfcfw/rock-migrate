module github.com/dfcfw/rock-migrate

go 1.24.0

require (
	github.com/robfig/cron/v3 v3.0.1
	github.com/xgfone/ship/v5 v5.3.1
	go.mongodb.org/mongo-driver v0.0.0-00010101000000-000000000000
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

// MongoDB 官方驱动自 1.10.0-beta1 及以后的版本，仅支持 MongoServer 6.0 及以上。
// 当前开发测试环境与线上环境低于 6.0，故将驱动版本锁定在 1.9.x。
//
// https://github.com/mongodb/mongo-go-driver/releases/tag/v1.10.0-beta1
//
// 驱动兼容查询：https://www.mongodb.com/zh-cn/docs/drivers/go/current/compatibility/
replace go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.9.4
