package flags

import (
	"fmt"
	"strings"

	"github.com/oherych/yeelightcli/internal/helper"

	"github.com/oherych/yeelight"
	"github.com/spf13/cobra"
)

const (
	FlagPowerMode = "power_mode"
)

var (
	powerModeValues = []powerModeOption{
		{Label: "default", Value: yeelight.PowerModeDefault},
		{Label: "color_temperature", Value: yeelight.PowerModeCT},
		{Label: "rgb", Value: yeelight.PowerModeRGB},
		{Label: "hsv", Value: yeelight.PowerModeHSV},
		{Label: "color_flow", Value: yeelight.PowerModeColorFlow},
		{Label: "color_night_light", Value: yeelight.PowerModeColorNightLight},
	}
)

type powerModeOption struct {
	Label string
	Value yeelight.PowerMode
}

func InjectPowerMode(cmd *cobra.Command) {
	labels := make([]string, len(powerModeValues))
	for i, option := range powerModeValues {
		labels[i] = option.Label
	}

	cmd.Flags().String(FlagPowerMode, "", fmt.Sprintf("Power mode (%s)", strings.Join(labels, ", ")))
}

func ReadPowerMode(cmd *cobra.Command) (yeelight.PowerMode, error) {
	val, err := cmd.Flags().GetString(FlagPowerMode)
	if err != nil {
		return 0, err
	}

	if val == "" {
		return yeelight.PowerModeDefault, nil
	}

	for _, option := range powerModeValues {
		if val == option.Label {
			return option.Value, nil
		}
	}

	return 0, helper.Error{Reason: "incorrect power mode name"}
}
