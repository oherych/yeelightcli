package arguments

import (
	"fmt"
	"strconv"

	"github.com/oherych/yeelight"
)

type SatArg struct{}

func (a SatArg) Name() string {
	return "sat"
}

func (a SatArg) Example() string {
	return ""
}

func (a SatArg) Description() string {
	return fmt.Sprintf("Saturation value in range %d - %d", yeelight.MinSat, yeelight.MaxSat)
}

func (a SatArg) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateSat(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}
