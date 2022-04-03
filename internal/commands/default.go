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

func (c DefaultCommand) Use() string {
	return "default"
}

func (c DefaultCommand) Short() string {
	return "Declare current light settings as default"
}

func (c DefaultCommand) Flags(cmd *cobra.Command) {
	flags.InjectBackground(cmd)
}

func (c DefaultCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c DefaultCommand) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.build(cmd, host).SetDefault(cmd.Context())
		},
		func() error {
			return c.build(cmd, host).SetBackgroundDefault(cmd.Context())
		},
	)
}
