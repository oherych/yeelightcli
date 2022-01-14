package commands

import (
	"github.com/oherych/yeelightcli/internal/flags"
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

func buildDefault(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "default", func(cmd *cobra.Command) {
		cmd.Short = "------"
		cmd.Args = cobra.NoArgs

		flags.InjectBackground(cmd)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			isBackground, err := flags.ReadBackground(cmd)
			if err != nil {
				return err
			}

			if isBackground {
				return build(cmd, args[0]).SetBackgroundDefault(cmd.Context())
			}

			return build(cmd, args[0]).SetDefault(cmd.Context())
		}
	})
}
