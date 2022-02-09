package power

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

type Root struct {
	Build helper.ClientBuilder
}

func (r Root) Use() string {
	return "power"
}

func (r Root) Short(cmd *cobra.Command) string {
	return "Switches the light on/off"
}

func (r Root) Long(cmd *cobra.Command) string {
	return ""
}

func (r Root) Flags(cmd *cobra.Command) {

}

func (r Root) SubCommand(cmd *cobra.Command) []helper.Command {
	return []helper.Command{
		OnOff{build: r.Build, isOn: true},
		OnOff{build: r.Build, isOn: false},
		PowerToggleCommand{build: r.Build},
		DevToggleCommand{build: r.Build},
	}
}

func (r Root) Args() []helper.Arg {
	return nil
}
