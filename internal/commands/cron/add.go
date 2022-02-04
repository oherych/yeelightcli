package cron

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Add struct {
	build helper.ClientBuilder
}

func (c Add) Use() string {
	return "add"
}

func (c Add) Short(cmd *cobra.Command) string {
	return "--"
}

func (c Add) Long(cmd *cobra.Command) string {
	return ""
}

func (c Add) Flags(cmd *cobra.Command) {
}

func (c Add) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c Add) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c Add) Run(cmd *cobra.Command, args []string) error {
	panic("implement me")
}
