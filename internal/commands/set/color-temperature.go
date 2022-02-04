package set

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
	"strconv"
)

type ColorTemperature struct {
	Build helper.ClientBuilder
}

func (c ColorTemperature) Use() string {
	return "ct"
}

func (c ColorTemperature) Short(cmd *cobra.Command) string {
	return "-----"
}

func (c ColorTemperature) Long(cmd *cobra.Command) string {
	return ""
}

func (c ColorTemperature) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c ColorTemperature) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c ColorTemperature) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.ColorTemperatureArg{}}
}

func (c ColorTemperature) Run(cmd *cobra.Command, args []string) error {
	effect, err := flags.ReadEffect(cmd)
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	colorTemperature, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return c.Build(cmd, args[0]).SetBackgroundColorTemperature(cmd.Context(), colorTemperature, effect, duration)
	}

	return c.Build(cmd, args[0]).SetColorTemperature(cmd.Context(), colorTemperature, effect, duration)
}
