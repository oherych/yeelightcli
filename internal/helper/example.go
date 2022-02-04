package helper

import (
	"github.com/spf13/cobra"
)

const (
	ExampleDomain = "192.168.1.79:55443"
)

func ExampleBuilder(cmd *cobra.Command, spec Command) string {
	result := "  " + cmd.CommandPath()

	for _, arg := range spec.Args() {
		result += " " + arg.Example()
	}

	return result
}
