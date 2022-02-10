package arguments

import (
	"github.com/oherych/yeelight"
	"strconv"
)

type Percentage struct{}

func (a Percentage) Name() string {
	return "percentage"
}

func (a Percentage) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateAdjustPercentage(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}

func (a Percentage) Example() string {
	return "75"
}
