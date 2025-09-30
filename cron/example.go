package cron

import (
	"github.com/zjutjh/jhgo/config"
	"github.com/zjutjh/jhgo/nlog"
)

type ExampleJob struct{}

func (ExampleJob) Run() {
	// 在此处编写定时任务业务逻辑
	nlog.Pick().WithField("app", config.AppName()).Debug("定时任务运行")
}
