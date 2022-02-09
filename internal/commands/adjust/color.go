package adjust

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Color struct {
	Build helper.ClientBuilder
}

func (c Color) Use() string {
	return "color"
}

func (c Color) Short(cmd *cobra.Command) string {
	return "-----"
}

func (c Color) Long(cmd *cobra.Command) string {
	return ""
}

func (c Color) Flags(cmd *cobra.Command) {
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c Color) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Percentage{}}
}

func (c Color) Run(cmd *cobra.Command, args []string) error {
	percentage, err := arguments.Percentage{}.Read(args[1])
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.Build(cmd, args[0]).AdjustColor(cmd.Context(), percentage, duration)
		},
		func() error {
			return c.Build(cmd, args[0]).AdjustBackgroundColor(cmd.Context(), percentage, duration)
		},
	)
}
