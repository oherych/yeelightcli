package adjust

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Bright struct {
	Build helper.ClientBuilder
}

func (c Bright) Use() string {
	return "brightness"
}

func (c Bright) Short(cmd *cobra.Command) string {
	return "-----"
}

func (c Bright) Long(cmd *cobra.Command) string {
	return ""
}

func (c Bright) Flags(cmd *cobra.Command) {
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c Bright) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Percentage{}}
}

func (c Bright) Run(cmd *cobra.Command, args []string) error {
	percentage, err := arguments.Percentage{}.Read(args[1])
	if err != nil {
		return err
	}

	duration, err := flags.ReadDurationFlag(cmd)
	if err != nil {
		return err
	}

	return flags.RunInBackground(cmd,
		func() error {
			return c.Build(cmd, args[0]).AdjustBright(cmd.Context(), percentage, duration)
		},
		func() error {
			return c.Build(cmd, args[0]).AdjustBackgroundBright(cmd.Context(), percentage, duration)
		},
	)
}
