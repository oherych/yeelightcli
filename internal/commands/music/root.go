package music

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Root struct {
	Build helper.ClientBuilder
}

func (c Root) Use() string {
	return "music"
}

func (c Root) Short() string {
	return "Start or stop music mode on a device"
}

func (c Root) Flags(cmd *cobra.Command) {}

func (c Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		On{Build: c.Build},
		Off{Build: c.Build},
	}
}

func (c Root) Args() []helper.Arg {
	return nil
}
