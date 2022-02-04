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

func (c OnOff) Short(cmd *cobra.Command) string {
	if c.isOn {
		return "Turn on device"
	}

	return "Turn off device"
}

func (c OnOff) Long(cmd *cobra.Command) string {
	return ""
}

func (c OnOff) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectPowerMode(cmd)
	flags.InjectBackground(cmd)
}

func (c OnOff) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c OnOff) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c OnOff) Run(cmd *cobra.Command, args []string) error {
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

	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return c.build(cmd, args[0]).BackgroundPower(cmd.Context(), c.isOn, powerMode, effect, duration)
	}

	return c.build(cmd, args[0]).Power(cmd.Context(), c.isOn, powerMode, effect, duration)
}
