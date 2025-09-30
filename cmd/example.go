package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zjutjh/jhgo/nlog"
)

func ExampleRun(cmd *cobra.Command, args []string) error {
	// 在此处编写命令业务逻辑
	nlog.Pick().WithFields(logrus.Fields{
		"cmd":  cmd.Name(),
		"args": args,
	}).Debug("命令运行")
	return nil
}
