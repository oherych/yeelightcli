package arguments

import (
	"strconv"
	"time"
)

type Minutes struct{}

func (a Minutes) Name() string {
	return "minutes"
}

func (a Minutes) Example() string {
	return (time.Minute).String()
}

func (a Minutes) Description() string {
	return "Timeout in minutes"
}

func (a Minutes) Read(in string) (time.Duration, error) {
	val, err := strconv.Atoi(in)
	if err != nil {
		return 0, standardError(a)
	}

	if val <= 0 {
		return 0, standardError(a)
	}

	return time.Duration(val) * time.Minute, err
}
