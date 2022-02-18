package set

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Brightness struct {
	Build helper.ClientBuilder
}

func (c Brightness) Use() string {
	return "brightness"
}

func (c Brightness) Short() string {
	return "Change brightness"
}

func (c Brightness) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c Brightness) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Bright{}}
}

func (c Brightness) Run(cmd *cobra.Command, args []string) error {
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

	bright, err := arguments.Bright{}.Read(args[1])
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.Build(cmd, host).SetBright(cmd.Context(), bright, effect, duration)
		},
		func() error {
			return c.Build(cmd, host).SetBackgroundBright(cmd.Context(), bright, effect, duration)
		},
	)
}
