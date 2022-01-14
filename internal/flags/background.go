package flags

import (
	"github.com/spf13/cobra"
)

const (
	flagBackground = "background"
)

func InjectBackground(cmd *cobra.Command) {
	cmd.Flags().Bool(flagBackground, false, "apply for background")
}

func ReadBackground(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(flagBackground)
}
