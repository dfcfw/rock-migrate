package cronfunc

import "github.com/dfcfw/rock-migrate/library/cronv3"

type CronInfo struct {
	Name string // 友好的名字
	Expr string // cron 表达式
	Func func() // 触发函数
}

func Add(crontab *cronv3.Crontab, infos []CronInfo) {
	for _, inf := range infos {
		crontab.AddJob(inf.Name, inf.Expr, inf.Func)
	}
}
