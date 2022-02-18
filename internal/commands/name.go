package commands

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

func (n Name) Short() string {
	// TODO: check
	return "Change device name"
}

func (n Name) Flags(cmd *cobra.Command) {

}

func (n Name) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.NameArg{}}
}

func (n Name) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	name, err := arguments.NameArg{}.Read(args[1])
	if err != nil {
		return err
	}

	return n.build(cmd, host).SetName(cmd.Context(), name)
}
