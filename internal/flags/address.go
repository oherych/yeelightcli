package flags

import "github.com/spf13/cobra"

const (
	flagAddress = "address"
)

func InjectAddress(cmd *cobra.Command) {
	cmd.Flags().StringP(flagAddress, "a", "", "device address")

	if err := cmd.MarkFlagRequired(flagAddress); err != nil {
		panic(err)
	}
}

func ReadAddress(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString(flagAddress)
}
