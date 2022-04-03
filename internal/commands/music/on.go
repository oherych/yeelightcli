package music

import (
	"github.com/oherych/yeelightcli/internal/arguments"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type On struct {
	Build helper.ClientBuilder
}

func (c On) Use() string {
	return "on"
}

func (c On) Short() string {
	return "Start listening audio server"
}

func (c On) Flags(cmd *cobra.Command) {
}

func (c On) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{}
}

func (c On) Args() []helper.Arg {
	return []helper.Arg{arguments.HostArg{}, arguments.MusicHost{}}
}

func (c On) Run(cmd *cobra.Command, args []string) error {
	host, err := arguments.HostArg{}.Read(args[0])
	if err != nil {
		return err
	}

	musicHost, musicPort, err := arguments.MusicHost{}.Read(args[1])
	if err != nil {
		return err
	}

	return c.Build(cmd, host).SetMusic(cmd.Context(), true, musicHost, musicPort)
}
