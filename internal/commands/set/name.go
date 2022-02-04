package set

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Name struct {
	build helper.ClientBuilder
}

func (n Name) Use() string {
	return "name"
}

func (n Name) Short(cmd *cobra.Command) string {
	return "Set device name"
}

func (n Name) Long(cmd *cobra.Command) string {
	return ""
}

func (n Name) Flags(cmd *cobra.Command) {

}

func (n Name) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.NameArg{}}
}

func (n Name) SubCommand(cmd *cobra.Command) []helper.Command {
	return nil
}

func (n Name) Run(cmd *cobra.Command, args []string) error {
	return n.build(cmd, args[0]).SetName(cmd.Context(), args[1])
}
