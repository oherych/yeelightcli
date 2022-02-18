package arguments

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"strings"
)

const (
	hostPort = "55443"
)

type HostArg struct{}

func (a HostArg) Name() string {
	return "host"
}

func (a HostArg) Example() string {
	return helper.ExampleDomain
}

func (a HostArg) Description() string {
	return "IP address or hostname of a device. If a port is missed then by default will be used " + hostPort + "."
}

func (a HostArg) Read(in string) (string, error) {
	if !strings.Contains(in, ":") {
		in += ":" + hostPort
	}

	return in, nil
}
