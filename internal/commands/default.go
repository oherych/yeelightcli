package commands

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type DefaultCommand struct {
	build helper.ClientBuilder
}

func (d DefaultCommand) Use() string {
	return "default"
}

func (d DefaultCommand) Short(cmd *cobra.Command) string {
	return "------"
}

func (d DefaultCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (d DefaultCommand) Flags(cmd *cobra.Command) {
	cmd.Args = cobra.NoArgs

	flags.InjectBackground(cmd)
}

func (d DefaultCommand) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (r DefaultCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (d DefaultCommand) Run(cmd *cobra.Command, args []string) error {
	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return d.build(cmd, args[0]).SetBackgroundDefault(cmd.Context())
	}

	return d.build(cmd, args[0]).SetDefault(cmd.Context())
}
