package arguments

import (
	"fmt"
	"github.com/oherych/yeelight"
	"strconv"
)

type AdjustPercentage struct{}

func (a AdjustPercentage) Name() string {
	return "percentage"
}

func (a AdjustPercentage) Description() string {
	return fmt.Sprintf("adjust percentage value in range %d - %d", yeelight.MinAdjustPercentage, yeelight.MaxAdjustPercentage)
}

func (a AdjustPercentage) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateAdjustPercentage(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}

func (a AdjustPercentage) Example() string {
	return "75"
}
