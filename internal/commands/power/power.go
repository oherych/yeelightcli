package power

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/spf13/cobra"
)

type OnOff struct {
	build helper.ClientBuilder
	isOn  bool
}

func (c OnOff) Use() string {
	if c.isOn {
		return "on"
	}

	return "off"
}

func (c OnOff) Short() string {
	if c.isOn {
		return "Switch the light on"
	}

	return "Switch the light off"
}

func (c OnOff) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectPowerMode(cmd)
	flags.InjectBackground(cmd)
}

func (c OnOff) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c OnOff) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	effect, err := flags.ReadEffect(cmd)
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	powerMode, err := flags.ReadPowerMode(cmd)
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.build(cmd, host).Power(cmd.Context(), c.isOn, powerMode, effect, duration)
		},
		func() error {
			return c.build(cmd, host).BackgroundPower(cmd.Context(), c.isOn, powerMode, effect, duration)
		},
	)
}
