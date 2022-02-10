package set

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Rgb struct {
	Build helper.ClientBuilder
}

func (r Rgb) Use() string {
	return "rgb"
}

func (r Rgb) Short(cmd *cobra.Command) string {
	return "Change color"
}

func (r Rgb) Long(cmd *cobra.Command) string {
	return ""
}

func (r Rgb) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (r Rgb) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Rgb{}}
}

func (r Rgb) Run(cmd *cobra.Command, args []string) error {
	color, err := arguments.Rgb{}.Read(args[1])
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
			return r.Build(cmd, args[0]).SetRGB(cmd.Context(), color, effect, duration)
		},
		func() error {
			return r.Build(cmd, args[0]).SetBackgroundRGB(cmd.Context(), color, effect, duration)
		},
	)
}
