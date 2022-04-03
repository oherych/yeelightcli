package flags

import (
	"github.com/oherych/yeelight"
	"time"

	"github.com/spf13/cobra"
)

const (
	FlagDurationName = "duration"
)

func InjectDurationFlag(cmd *cobra.Command) {
	cmd.Flags().DurationP(FlagDurationName, "d", yeelight.MinDuration, "Effect duration")
}

func ReadDurationFlag(cmd *cobra.Command) (time.Duration, error) {
	val, err := cmd.Flags().GetDuration(FlagDurationName)
	if err != nil {
		return 0, err
	}

	return val, nil
}
