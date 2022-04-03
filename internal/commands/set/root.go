package set

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Root struct {
	Build helper.ClientBuilder
}

func (r Root) Use() string {
	return "set"
}

func (r Root) Short() string {
	return "Specify the light parameter"
}

func (r Root) Flags(cmd *cobra.Command) {}

func (r Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		Rgb{Build: r.Build},
		Brightness{Build: r.Build},
		ColorTemperature{Build: r.Build},
		Hsv{Build: r.Build},
	}
}

func (r Root) Args() []helper.Arg {
	return nil
}
