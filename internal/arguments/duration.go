package arguments

import (
	"time"
)

type Duration struct{}

func (a Duration) Name() string {
	return "duration"
}

func (a Duration) Example() string {
	return (time.Minute).String()
}

func (a Duration) Read(in string) (time.Duration, error) {
	val, err := time.ParseDuration(in)
	if err != nil {
		return 0, standardError(a)
	}

	return val, err
}
