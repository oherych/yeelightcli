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

func (c Color) Short() string {
	return "Change color value"
}

func (c Color) Flags(cmd *cobra.Command) {
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c Color) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.AdjustPercentage{}}
}

func (c Color) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	percentage, err := arguments.AdjustPercentage{}.Read(args[1])
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.Build(cmd, host).AdjustColor(cmd.Context(), percentage, duration)
		},
		func() error {
			return c.Build(cmd, host).AdjustBackgroundColor(cmd.Context(), percentage, duration)
		},
	)
}
