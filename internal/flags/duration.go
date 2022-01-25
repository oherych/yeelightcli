package flags

import (
	"time"

	"github.com/spf13/cobra"
)

const (
	FlagDurationName     = "duration"
	MinimalDurationValue = 30 * time.Millisecond
)

func InjectDurationFlag(cmd *cobra.Command) {
	cmd.Flags().DurationP(FlagDurationName, "d", MinimalDurationValue, "Effect duration")
}

func ReadDurationFlag(cmd *cobra.Command) (time.Duration, error) {
	val, err := cmd.Flags().GetDuration(FlagDurationName)
	if err != nil {
		return 0, err
	}

	return val, nil
}
