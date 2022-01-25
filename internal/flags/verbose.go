package flags

import "github.com/spf13/cobra"

const (
	FlagVerboseName = "verbose"
)

func InjectVerboseFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(FlagVerboseName, "v", false, "verbose output")
}

func ReadVerboseFlag(cmd *cobra.Command) (bool, error) {
	return cmd.Flags().GetBool(FlagVerboseName)
}
