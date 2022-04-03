package power

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type ToggleCommand struct {
	build helper.ClientBuilder
}

func (c ToggleCommand) Use() string {
	return "toggle"
}

func (c ToggleCommand) Short() string {
	return "Toggle the light"
}

func (c ToggleCommand) Flags(cmd *cobra.Command) {
	flags.InjectBackground(cmd)
}

func (c ToggleCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c ToggleCommand) Run(cmd *cobra.Command, args []string) error {
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
