package commands

import (
	"fmt"

	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/spf13/cobra"
)

func buildPowerOnOff(parent *cobra.Command, build clientBuilder, use string, short string, on bool) {
	helper.BuildCommand(parent, use+" [host]", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s", cmd.CommandPath(), exampleDomain)
		cmd.Short = short

		flags.InjectEffect(cmd)
		flags.InjectDurationFlag(cmd)
		flags.InjectPowerMode(cmd)
		flags.InjectBackground(cmd)
		flags.InjectAddress(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			address, err := flags.ReadAddress(cmd)
			if err != nil {
				return err
			}

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

			isBackground, err := flags.ReadBackground(cmd)
			if err != nil {
				return err
			}

			if isBackground {
				return build(cmd, address).BackgroundPower(cmd.Context(), on, powerMode, effect, duration)
			}

			return build(cmd, address).Power(cmd.Context(), on, powerMode, effect, duration)
		}
	})

}

func buildPowerToggle(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "toggle", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s", cmd.CommandPath(), exampleDomain)
		cmd.Short = "Toggle power"
		cmd.Args = cobra.ExactValidArgs(1)

		flags.InjectBackground(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			isBackground, err := flags.ReadBackground(cmd)
			if err != nil {
				return err
			}

			if isBackground {
				return build(cmd, args[0]).BackgroundToggle(cmd.Context())
			}

			return build(cmd, args[0]).Toggle(cmd.Context())
		}
	})
}

func buildDevToggle(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "device_toggle", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s", cmd.CommandPath(), exampleDomain)
		cmd.Short = "------"
		cmd.Args = cobra.ExactValidArgs(1)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return build(cmd, args[0]).DevToggle(cmd.Context())
		}
	})
}
