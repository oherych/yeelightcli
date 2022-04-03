package music

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Off struct {
	Build helper.ClientBuilder
}

func (c Off) Use() string {
	return "off"
}

func (c Off) Short() string {
	return "Stop music"
}

func (c Off) Flags(cmd *cobra.Command) {
}

func (c Off) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{}
}

func (c Off) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}}
}

func (c Off) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	return c.Build(cmd, host).SetMusic(cmd.Context(), false, "", 0)
}
