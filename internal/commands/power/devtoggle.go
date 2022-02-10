package power

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type DevToggleCommand struct {
	build helper.ClientBuilder
}

func (c DevToggleCommand) Use() string {
	return "dev-toggle"
}

func (c DevToggleCommand) Short(cmd *cobra.Command) string {
	return "Toggle both the main and background light at the same time"
}

func (c DevToggleCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (c DevToggleCommand) Flags(cmd *cobra.Command) {

}

func (c DevToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c DevToggleCommand) Run(cmd *cobra.Command, args []string) error {
	return c.build(cmd, args[0]).DevToggle(cmd.Context())
}
