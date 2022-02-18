package arguments

import (
	"github.com/oherych/yeelight"
	"strconv"
)

type HueArg struct{}

func (a HueArg) Name() string {
	return "hue"
}

func (a HueArg) Example() string {
	return ""
}

func (a HueArg) Description() string {
	// TODO: add text
	return ""
}

func (a HueArg) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateHue(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}
