package power

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type PowerToggleCommand struct {
	build helper.ClientBuilder
}

func (p PowerToggleCommand) Use() string {
	return "toggle"
}

func (p PowerToggleCommand) Short(cmd *cobra.Command) string {
	return "Toggle power"
}

func (p PowerToggleCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (p PowerToggleCommand) Flags(cmd *cobra.Command) {
	flags.InjectBackground(cmd)
}

func (p PowerToggleCommand) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (r PowerToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (p PowerToggleCommand) Run(cmd *cobra.Command, args []string) error {
	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return p.build(cmd, args[0]).BackgroundToggle(cmd.Context())
	}

	return p.build(cmd, args[0]).Toggle(cmd.Context())
}
