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

func (r Root) Short(cmd *cobra.Command) string {
	return "-----"
}

func (r Root) Long(cmd *cobra.Command) string {
	return ""
}

func (r Root) Flags(cmd *cobra.Command) {}

func (r Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		Name{build: r.Build},
		Rgb{Build: r.Build},
		Brightness{Build: r.Build},
		ColorTemperature{Build: r.Build},
		Hsv{Build: r.Build},
	}
}

func (r Root) Args() []helper.Arg {
	return nil
}
