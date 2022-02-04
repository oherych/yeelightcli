package set

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
	return "bright"
}

func (c Bright) Short(cmd *cobra.Command) string {
	return "Set the brightness"
}

func (c Bright) Long(cmd *cobra.Command) string {
	return ""
}

func (c Bright) Flags(cmd *cobra.Command) {
	flags.InjectEffect(cmd)
	flags.InjectDurationFlag(cmd)
	flags.InjectBackground(cmd)
}

func (c Bright) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c Bright) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.Bright{}}
}

func (c Bright) Run(cmd *cobra.Command, args []string) error {
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

	isBackground, err := flags.ReadBackground(cmd)
	if err != nil {
		return err
	}

	if isBackground {
		return c.Build(cmd, args[0]).SetBackgroundBright(cmd.Context(), bright, effect, duration)
	}

	return c.Build(cmd, args[0]).SetBright(cmd.Context(), bright, effect, duration)
}
