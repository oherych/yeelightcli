package set

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Hsv struct {
	Build helper.ClientBuilder
}

func (h Hsv) Use() string {
	return "hsv"
}

func (h Hsv) Short(cmd *cobra.Command) string {
	return "----"
}

func (h Hsv) Long(cmd *cobra.Command) string {
	return ""
}

func (h Hsv) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (h Hsv) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (h Hsv) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.HsvArg{}}
}

func (h Hsv) Run(cmd *cobra.Command, args []string) error {
	effect, err := flags.ReadEffect(cmd)
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	// TODO

	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return h.Build(cmd, args[0]).SetBackgroundHSV(cmd.Context(), 0, 0, effect, duration)
	}

	return h.Build(cmd, args[0]).SetHSV(cmd.Context(), 0, 0, effect, duration)
}
