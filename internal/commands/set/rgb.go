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

func (r Rgb) Short() string {
	return "Change RGB"
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
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

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
			return r.Build(cmd, host).SetRGB(cmd.Context(), color, effect, duration)
		},
		func() error {
			return r.Build(cmd, host).SetBackgroundRGB(cmd.Context(), color, effect, duration)
		},
	)
}
