package cronfunc

type CronInfo struct {
	Name string // 友好的名字
	Expr string // cron 表达式
	Func func() // 触发函数
}

func Register() {
}
