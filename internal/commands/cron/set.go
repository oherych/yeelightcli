package cron

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Add struct {
	build helper.ClientBuilder
}

func (c Add) Use() string {
	return "set"
}

func (c Add) Short() string {
	return "Define timeout"
}

func (c Add) Flags(cmd *cobra.Command) {
}

func (c Add) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.OffOn{}, arguments.Minutes{}}
}

func (c Add) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	on, err := arguments.OffOn{}.Read(args[1])
	if err != nil {
		return err
	}

	duration, err := arguments.Minutes{}.Read(args[2])
	if err != nil {
		return err
	}

	return c.build(cmd, host).AddCron(cmd.Context(), on, duration)
}
