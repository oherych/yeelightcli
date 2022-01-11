package flags

import (
	"fmt"
	"strings"

	"github.com/oherych/yeelight"
	"github.com/spf13/cobra"
)

const (
	FlagEffectName = "effect"
)

var (
	effectValues = []effectOption{
		{Label: "sudden", Value: yeelight.EffectSudden},
		{Label: "smooth", Value: yeelight.EffectSmooth},
	}
)

type effectOption struct {
	Label string
	Value yeelight.Effect
}

func InjectEffect(cmd *cobra.Command) {
	labels := make([]string, len(effectValues))
	for i, option := range effectValues {
		labels[i] = option.Label
	}

	cmd.Flags().String(FlagEffectName, string(yeelight.EffectSmooth), fmt.Sprintf("Effect name (%s)", strings.Join(labels, ", ")))
}

func ReadEffect(cmd *cobra.Command) (yeelight.Effect, error) {
	val, err := cmd.Flags().GetString(FlagEffectName)
	if err != nil {
		return "", err
	}

	for _, option := range effectValues {
		if val == option.Label {
			return option.Value, nil
		}
	}

	panic("TODO")
}
