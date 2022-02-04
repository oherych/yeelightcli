package cron

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Delete struct {
	build helper.ClientBuilder
}

func (c Delete) Use() string {
	return "delete"
}

func (c Delete) Short(cmd *cobra.Command) string {
	return ""
}

func (c Delete) Long(cmd *cobra.Command) string {
	return ""
}

func (c Delete) Flags(cmd *cobra.Command) {

}

func (c Delete) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c Delete) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c Delete) Run(cmd *cobra.Command, args []string) error {
	return c.build(cmd, args[0]).DeleteCron(cmd.Context(), false)
}
