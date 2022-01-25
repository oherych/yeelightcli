package helper

import "github.com/spf13/cobra"

func BuildCommand(parent *cobra.Command, use string, fn func(cmd *cobra.Command)) {
	c := &cobra.Command{Use: use}
	c.Flags().SortFlags = false

	parent.AddCommand(c)

	fn(c)
}
