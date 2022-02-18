package adjust

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Root struct {
	Build helper.ClientBuilder
}

func (c Root) Use() string {
	return "adjust"
}

func (c Root) Short() string {
	return "Adjust the light parameter"
}

func (c Root) Flags(cmd *cobra.Command) {
}

func (c Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		Bright{Build: c.Build},
		ColorTemperature{Build: c.Build},
		Color{Build: c.Build},
	}
}

func (c Root) Args() []helper.Arg {
	return nil
}
