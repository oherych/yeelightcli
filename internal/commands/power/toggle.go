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

func (c PowerToggleCommand) Use() string {
	return "toggle"
}

func (c PowerToggleCommand) Short(cmd *cobra.Command) string {
	return "Toggle the light"
}

func (c PowerToggleCommand) Long(cmd *cobra.Command) string {
	return ""
}

func (c PowerToggleCommand) Flags(cmd *cobra.Command) {
	flags.InjectBackground(cmd)
}

func (c PowerToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c PowerToggleCommand) Run(cmd *cobra.Command, args []string) error {
	return flags.RunInBackground(cmd,
		func() error {
			return c.build(cmd, args[0]).Toggle(cmd.Context())
		},
		func() error {
			return c.build(cmd, args[0]).BackgroundToggle(cmd.Context())
		},
	)
}
