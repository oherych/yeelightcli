package arguments

import (
	"strconv"
)

type Rgb struct{}

func (a Rgb) Name() string {
	return "rgb"
}

func (a Rgb) Description() string {
	// TODO: add text
	return ""
}

func (a Rgb) Read(in string) (int, error) {
	val, err := strconv.ParseUint(in, 16, 32)
	if err != nil || val <= 0 {
		return 0, standardError(a)
	}

	return int(val), err
}

func (a Rgb) Example() string {
	return "FF0000"
}
