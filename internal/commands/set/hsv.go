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
	return "Change hue and saturation"
}

func (h Hsv) Long(cmd *cobra.Command) string {
	return ""
}

func (h Hsv) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (h Hsv) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.HueArg{}, arguments.SatArg{}}
}

func (h Hsv) Run(cmd *cobra.Command, args []string) error {
	hue, err := arguments.HueArg{}.Read(args[1])
	if err != nil {
		return err
	}

	sat, err := arguments.SatArg{}.Read(args[1])
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

	return flags.RunInBackground(cmd,
		func() error {
			return h.Build(cmd, args[0]).SetHSV(cmd.Context(), hue, sat, effect, duration)
		},
		func() error {
			return h.Build(cmd, args[0]).SetBackgroundHSV(cmd.Context(), hue, sat, effect, duration)
		},
	)
}
