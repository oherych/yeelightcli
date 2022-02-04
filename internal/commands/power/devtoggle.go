package power

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type DevToggleCommand struct {
	build helper.ClientBuilder
}

func (d DevToggleCommand) Use() string {
	return "dev-toggle"
}

func (d DevToggleCommand) Short(cmd *cobra.Command) string {
	return "------"
}

func (d DevToggleCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (d DevToggleCommand) Flags(cmd *cobra.Command) {

}

func (d DevToggleCommand) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (r DevToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (d DevToggleCommand) Run(cmd *cobra.Command, args []string) error {
	return d.build(cmd, args[0]).DevToggle(cmd.Context())
}
