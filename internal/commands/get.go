package commands

import (
	"fmt"
	"strings"

	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/spf13/cobra"
)

func buildGet(parent *cobra.Command, build clientBuilder) {
	helper.BuildCommand(parent, "get [host] [parameter1, parameter2..]", func(cmd *cobra.Command) {
		cmd.Short = "Read device values"
		cmd.Args = cobra.ExactArgs(2)

		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			properties := strings.Split(args[1], ",")
			for i := range properties {
				properties[i] = strings.TrimSpace(properties[i])
			}

			result, err := build(cmd, args[0]).GetProperties(cmd.Context(), properties)
			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		}
	})
}
