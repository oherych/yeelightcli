package arguments

import (
	"fmt"
	"github.com/oherych/yeelight"
	"strconv"
)

type ColorTemperatureArg struct{}

func (a ColorTemperatureArg) Name() string {
	return "color temperature"
}

func (a ColorTemperatureArg) Example() string {
	return "6500"
}

func (a ColorTemperatureArg) Description() string {
	// TODO: add text
	return fmt.Sprintf("Color temperature value [%d - %d]", yeelight.MinColorTemperature, yeelight.MaxColorTemperature)
}

func (a ColorTemperatureArg) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateColorTemperature(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}
