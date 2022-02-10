package arguments

import (
	"github.com/oherych/yeelightcli/internal/helper"
)

type HostArg struct{}

func (a HostArg) Name() string {
	return "host"
}

func (a HostArg) Example() string {
	return helper.ExampleDomain
}

func (a HostArg) Read(in string) (string, error) {
	return in, nil
}
