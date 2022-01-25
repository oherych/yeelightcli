package flags

import (
	"errors"
	"github.com/spf13/cobra"
)

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

func HostsArg() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("[host] is required")
		}

		return nil
	}
}
