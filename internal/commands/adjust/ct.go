package adjust

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type ColorTemperature struct {
	Build helper.ClientBuilder
}

func (c ColorTemperature) Use() string {
	return "ct"
}

func (c ColorTemperature) Short() string {
	return "Change color temperature"
}

func (c ColorTemperature) Flags(cmd *cobra.Command) {
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c ColorTemperature) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.AdjustPercentage{}}
}

func (c ColorTemperature) Run(cmd *cobra.Command, args []string) error {
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
			return c.Build(cmd, host).AdjustColorTemperature(cmd.Context(), percentage, duration)
		},
		func() error {
			return c.Build(cmd, host).AdjustBackgroundColorTemperature(cmd.Context(), percentage, duration)
		},
	)
}
