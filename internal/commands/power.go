package commands

import (
	"fmt"

	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/spf13/cobra"
)

func buildPower(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "power", func(cmd *cobra.Command) {
		cmd.Short = "Power control"

		buildPowerOnOff(cmd, build, "on", "Turn on device", true)
		buildPowerOnOff(cmd, build, "off", "Turn off device", false)
		buildPowerToggle(cmd, build)
	})
}

func buildPowerOnOff(parent *cobra.Command, build clientBuilder, use string, short string, on bool) {
	helper.BuildCommand(parent, use+" [host]", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s", cmd.CommandPath(), exampleDomain)
		cmd.Short = short
		cmd.Args = cobra.ExactValidArgs(1)

		flags.InjectEffect(cmd)
		flags.InjectDurationFlag(cmd)
		flags.InjectPowerMode(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			effect, err := flags.ReadEffect(cmd)
			if err != nil {
				return err
			}

			duration, err := flags.ReadDurationFlag(cmd)
			if err != nil {
				return err
			}

			powerMode, err := flags.ReadPowerMode(cmd)
			if err != nil {
				return err
			}

			return build(cmd, args[0]).Power(cmd.Context(), on, powerMode, effect, duration)
		}
	})

}

func buildPowerToggle(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "toggle", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s", cmd.CommandPath(), exampleDomain)
		cmd.Short = "Toggle power"
		cmd.Args = cobra.ExactValidArgs(1)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return build(cmd, args[0]).Toggle(cmd.Context())
		}
	})
}
