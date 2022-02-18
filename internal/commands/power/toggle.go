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

func (c PowerToggleCommand) Short() string {
	// TODO: check
	return "Toggle the light"
}

func (c PowerToggleCommand) Flags(cmd *cobra.Command) {
	flags.InjectBackground(cmd)
}

func (c PowerToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c PowerToggleCommand) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.build(cmd, host).Toggle(cmd.Context())
		},
		func() error {
			return c.build(cmd, host).BackgroundToggle(cmd.Context())
		},
	)
}
