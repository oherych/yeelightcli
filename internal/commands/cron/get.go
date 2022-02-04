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

func (c Get) Short(cmd *cobra.Command) string {
	return "Get current setting"
}

func (c Get) Long(cmd *cobra.Command) string {
	return ""
}

func (c Get) Flags(cmd *cobra.Command) {

}

func (c Get) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (c Get) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c Get) Run(cmd *cobra.Command, args []string) error {
	d, err := c.build(cmd, args[0]).GetCron(cmd.Context(), false)
	if err != nil {
		return err
	}

	fmt.Println(d)

	return nil
}
