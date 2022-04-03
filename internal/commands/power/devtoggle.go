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

func (c DevToggleCommand) Short() string {
	return "Toggle both the main and background light at the same time"
}

func (c DevToggleCommand) Flags(cmd *cobra.Command) {

}

func (c DevToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c DevToggleCommand) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	return c.build(cmd, host).DevToggle(cmd.Context())
}
