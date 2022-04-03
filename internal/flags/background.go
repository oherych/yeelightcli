package flags

import (
	"github.com/spf13/cobra"
)

const (
	flagBackground = "background"
)

func InjectBackground(cmd *cobra.Command) {
	cmd.Flags().BoolP(flagBackground, "b", false, "apply for background")
}

func RunInBackground(cmd *cobra.Command, main, background func() error) error {
	isBackground, err := cmd.Flags().GetBool(flagBackground)
	if err != nil {
		return err
	}

	if isBackground {
		return background()
	}

	return main()
}
