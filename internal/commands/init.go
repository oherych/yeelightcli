package commands

import (
	"github.com/oherych/yeelightcli/internal/helper"
	"github.com/spf13/cobra"
)

func Root(applicationName string, build helper.ClientBuilder, discovery helper.DiscoveryFn) *cobra.Command {
	cobra.EnableCommandSorting = false

	command := &cobra.Command{
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	initCommand(command, RootCommand{
		applicationName: applicationName,
		build:           build,
		discovery:       discovery,
	})

	return command
}

func initCommand(base *cobra.Command, spec helper.Command) {
	base.SilenceUsage = true
	base.Use = buildUsage(spec)
	base.Short = spec.Short(base)
	base.Long = spec.Long(base)

	if imp, ok := spec.(helper.CommandExample); ok {
		base.Example = imp.Example(base)
	} else if _, ok := spec.(helper.CommandRun); ok {
		base.Example = helper.ExampleBuilder(base, spec)
	}

	if imp, ok := spec.(helper.CommandRun); ok {
		base.RunE = func(cmd *cobra.Command, args []string) error {
			if len(spec.Args()) != len(args) {
				return helper.Error{
					Reason:      "wrong number of arguments",
					Instruction: "\n Examples:\n" + cmd.Example,
				}
			}

			return imp.Run(cmd, args)
		}
	}

	spec.Flags(base)

	if imp, ok := spec.(helper.CommandSubCommand); ok {
		for _, sub := range imp.SubCommand(base) {
			cmd := &cobra.Command{}
			base.AddCommand(cmd)

			initCommand(cmd, sub)
		}
	}

}

func buildUsage(spec helper.Command) string {
	result := spec.Use()

	for _, arg := range spec.Args() {
		result += " [" + arg.Name() + "]"
	}

	return result
}
