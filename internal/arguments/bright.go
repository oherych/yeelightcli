package arguments

import (
	"github.com/oherych/yeelight"
	"strconv"
)

type Bright struct{}

func (a Bright) Name() string {
	return "bright"
}

func (a Bright) Example() string {
	return "10"
}

func (a Bright) Read(in string) (int, error) {
	val, err := strconv.Atoi(in)
	if err != nil || yeelight.ValidateBright(val) != nil {
		return 0, standardError(a)
	}

	return val, err
}
