package commands

import (
	"fmt"
	"strconv"

	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

func buildSetColor(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "rgb", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s red", cmd.CommandPath(), exampleDomain)
		cmd.Short = "Set color of lamp"
		cmd.Args = cobra.ExactValidArgs(2)

		flags.InjectEffect(cmd)
		flags.InjectDurationFlag(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			effect, err := flags.ReadEffect(cmd)
			if err != nil {
				return err
			}

			duration, err := flags.ReadDurationFlag(cmd)
			if err != nil {
				return err
			}

			color, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			return build(cmd, args[0]).SetRGB(cmd.Context(), color, effect, duration)
		}

	})
}

func buildSetBright(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "bright", func(cmd *cobra.Command) {
		cmd.Example = fmt.Sprintf("  %s %s red", cmd.CommandPath(), exampleDomain)
		cmd.Short = "-----"
		cmd.Args = cobra.ExactValidArgs(2)

		flags.InjectEffect(cmd)
		flags.InjectDurationFlag(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			effect, err := flags.ReadEffect(cmd)
			if err != nil {
				return err
			}

			duration, err := flags.ReadDurationFlag(cmd)
			if err != nil {
				return err
			}

			color, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			return build(cmd, args[0]).SetBright(cmd.Context(), color, effect, duration)
		}

	})
}

func buildSetName(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "name", func(cmd *cobra.Command) {
		cmd.Example = ""
		cmd.Short = "Set device name"
		cmd.Args = cobra.ExactValidArgs(2)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return build(cmd, args[0]).SetName(cmd.Context(), args[1])
		}
	})
}
