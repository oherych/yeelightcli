package cron

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Delete struct {
	build helper.ClientBuilder
}

func (c Delete) Use() string {
	return "delete"
}

func (c Delete) Short(cmd *cobra.Command) string {
	return ""
}

func (c Delete) Long(cmd *cobra.Command) string {
	return ""
}

func (c Delete) Flags(cmd *cobra.Command) {

}

func (c Delete) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.OffOn{}}
}

func (c Delete) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	on, err := arguments.OffOn{}.Read(args[1])
	if err != nil {
		return err
	}

	return c.build(cmd, host).DeleteCron(cmd.Context(), on)
}
