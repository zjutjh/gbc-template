package register

import (
	"github.com/spf13/cobra"
	"github.com/zjutjh/jhgo/foundation/command"
	"github.com/zjutjh/jhgo/foundation/crontab"
	"github.com/zjutjh/jhgo/foundation/httpserver"
)

func Command(root *cobra.Command) {
	// 基础命令
	command.Add("server", httpserver.CommandRegister(Route))
	command.Add("cron", crontab.CommandRegister(Cron))

	// 业务命令
	//command.Add("xxx", cmd.XXXRun)
}
