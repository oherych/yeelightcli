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
	flags.InjectBackground(cmd)
}

func (r DefaultCommand) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (d DefaultCommand) Run(cmd *cobra.Command, args []string) error {
	return flags.RunInBackground(cmd,
		func() error {
			return d.build(cmd, args[0]).SetDefault(cmd.Context())
		},
		func() error {
			return d.build(cmd, args[0]).SetBackgroundDefault(cmd.Context())
		},
	)
}
