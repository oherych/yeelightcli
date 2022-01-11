package commands

import (
	"fmt"

	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

func buildCron(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "cron", func(cmd *cobra.Command) {
		cmd.Short = "Timeout management"

		buildCronGet(cmd, build)
		buildCronAdd(cmd, build)
		buildCronDelete(cmd, build)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		}
	})

}

func buildCronGet(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "get [host]", func(cmd *cobra.Command) {
		cmd.Short = "Get current setting"

		cmd.Args = cobra.ExactValidArgs(1)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			d, err := build(cmd, args[0]).GetCron(cmd.Context(), false)
			if err != nil {
				return err
			}

			fmt.Println(d)

			return nil
		}
	})
}

func buildCronAdd(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "add [host] [timeout in seconds]", func(cmd *cobra.Command) {
		cmd.Short = "Set timeout setting"

		cmd.Args = cobra.ExactValidArgs(2)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return nil
		}
	})
}

func buildCronDelete(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "delete [host]", func(cmd *cobra.Command) {
		cmd.Short = "Delete setting"

		cmd.Args = cobra.ExactValidArgs(1)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			return build(cmd, args[0]).DeleteCron(cmd.Context(), false)
		}
	})
}
