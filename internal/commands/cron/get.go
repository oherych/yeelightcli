package cron

import (
	"fmt"

	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/spf13/cobra"
)

type Get struct {
	build helper.ClientBuilder
}

func (c Get) Use() string {
	return "get"
}

func (c Get) Short() string {
	// TODO: check
	return "Get current setting"
}

func (c Get) Flags(cmd *cobra.Command) {

}

func (c Get) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.OffOn{}}
}

func (c Get) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	on, err := arguments.OffOn{}.Read(args[1])
	if err != nil {
		return err
	}

	d, err := c.build(cmd, host).GetCron(cmd.Context(), on)
	if err != nil {
		return err
	}

	fmt.Println(d)

	return nil
}
